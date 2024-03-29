package app

import (
	"context"
	"gitea.com/lzhuk/forum/internal/server"
	"gitea.com/lzhuk/forum/pkg/config"
	"log"
	"net/http"
)

type Server struct {
	ServerHTTP *http.Server
}

func NewServer(cfg config.Config, r server.Router) *Server {
	return &Server{
		ServerHTTP: &http.Server{
			Addr:    ":" + cfg.Port,
			Handler: r,
		},
	}
}

func (s *Server) Start(ctx context.Context) error {
	go func() { // parallel go func() int {return -1} ()
		if err := s.ServerHTTP.ListenAndServe(); err != nil {
			log.Println(err)
			return
		}
	}()
	<-ctx.Done()
	return s.ServerHTTP.Shutdown(context.Background())
}

func (s *Server) Shutdown(ctx context.Context) error {
	return s.ServerHTTP.Shutdown(ctx)
}
