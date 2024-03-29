package main

import (
	"context"
	"database/sql"
	"fmt"
	"gitea.com/lzhuk/forum/internal/app"
	"gitea.com/lzhuk/forum/internal/service/comment"
	"gitea.com/lzhuk/forum/internal/service/post"
	"gitea.com/lzhuk/forum/internal/service/user"
	"log"
	"os/signal"
	"syscall"
	"time"

	"gitea.com/lzhuk/forum/internal/repository"
	"gitea.com/lzhuk/forum/internal/server"
	"gitea.com/lzhuk/forum/internal/service"
	"gitea.com/lzhuk/forum/pkg/config"
	_ "github.com/mattn/go-sqlite3"
)

func main() {
	// Загрузка конфигурации
	cfg, err := config.Load()
	if err != nil {
		log.Fatal(err)
	}

	// Инициализация родительского контекста
	_, cancel := context.WithCancel(context.Background())
	defer cancel()

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

	usersRepo := repository.NewUserRepo(db)
	sessionRepo := repository.NewSessionRepository(db)
	postsRepo := repository.NewPostsRepo(db)
	commentsRepo := repository.NewCommentsRepo(db)

	usersService := user.NewUserService(usersRepo)
	sessionsService := user.NewSessionService(sessionRepo)
	postsService := post.NewPostsService(postsRepo)
	commentsService := comment.NewCommentsService(commentsRepo)

	services := service.NewService(usersService, sessionsService, postsService, commentsService)
	handler := server.NewHandler(services)
	router := server.NewRouter(&handler)
	s := app.NewServer(cfg, router)

	// ???
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

	//err = s.ListenAndServe()
	//if err != nil {
	//	log.Println("Сервер на порту %s не запущен. Ошибка: %s", cfg.Port, err)
	//}
}
