package main

import (
	"context"
	"database/sql"
	"fmt"
	"log"
	"os"
	"os/signal"
	"syscall"
	"time"

	comment2 "gitea.com/lzhuk/forum/internal/repository/comment"
	post2 "gitea.com/lzhuk/forum/internal/repository/post"
	user2 "gitea.com/lzhuk/forum/internal/repository/user"

	"gitea.com/lzhuk/forum/internal/app"
	"gitea.com/lzhuk/forum/internal/service/comment"
	"gitea.com/lzhuk/forum/internal/service/post"
	"gitea.com/lzhuk/forum/internal/service/user"

	"gitea.com/lzhuk/forum/internal/server"
	"gitea.com/lzhuk/forum/internal/service"
	"gitea.com/lzhuk/forum/pkg/config"
	_ "github.com/mattn/go-sqlite3"
)

func main() {
	logFile, err := os.Create("./internal/logs/result.log")
	if err != nil {
		log.Println("Doesn't open file: ", err, ".")
	}
	defer logFile.Close()
	log.SetOutput(logFile)
	log.SetPrefix("Log: ")
	log.SetFlags(log.Ldate | log.Ltime | log.Lshortfile)

	cfg, err := config.Load()
	if err != nil {
		log.Fatal(err)
	}

	db, err := sql.Open("sqlite3", "forum.sqlite3")
	if err != nil {
		fmt.Println(err)
	}

	defer func() {
		err := db.Close()
		if err != nil {
			log.Println(err)
		}
	}()

	usersRepo := user2.NewUserRepo(db)
	sessionRepo := user2.NewSessionRepository(db)
	postsRepo := post2.NewPostsRepo(db)
	likePostRepo := post2.NewLikePostRepository(db)
	commentsRepo := comment2.NewCommentsRepo(db)
	likecommentsRepo := comment2.NewLikeCommentRepository(db)

	usersService := user.NewUserService(usersRepo)
	sessionsService := user.NewSessionService(sessionRepo)
	postsService := post.NewPostsService(postsRepo)
	likePostsService := post.NewLikePostService(likePostRepo)
	commentsService := comment.NewCommentsService(commentsRepo)
	likecommentsService := comment.NewLikeCommentService(likecommentsRepo)

	services := service.NewService(usersService, sessionsService, postsService, commentsService, likePostsService, likecommentsService)
	handler := server.NewHandler(services)
	router := server.NewRouter(&handler)
	s := app.NewServer(cfg, router)

	ctx, stop := signal.NotifyContext(context.Background(), syscall.SIGINT, syscall.SIGTERM)
	defer stop()

	go func() {
		if err := s.Start(ctx); err != nil {
			log.Println(err)
			return
		}
	}()

	<-ctx.Done()

	shutdown, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	if err != s.Shutdown(shutdown) {
		log.Println(err)
		return
	}
}
