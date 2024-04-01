package main

import (
	"context"
	"database/sql"
	"fmt"
	"log"
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
	commentsRepo := comment2.NewCommentsRepo(db)

	usersService := user.NewUserService(usersRepo)
	sessionsService := user.NewSessionService(sessionRepo)
	postsService := post.NewPostsService(postsRepo)
	commentsService := comment.NewCommentsService(commentsRepo)

	services := service.NewService(usersService, sessionsService, postsService, commentsService)
	handler := server.NewHandler(services)
	router := server.NewRouter(&handler)
	s := app.NewServer(cfg, router)

	// It is work but it need for creating context 1
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
