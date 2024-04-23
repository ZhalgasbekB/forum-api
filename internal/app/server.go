package app

import (
	"context"
	"log"
	"net/http"

	"gitea.com/lzhuk/forum/internal/server"
	"gitea.com/lzhuk/forum/pkg/config"
)

type Server struct {
	ServerHTTP *http.Server
}

func NewServer(cfg config.Config, r server.Router) *Server {
	//tlsConfig := &tls.Config{
	//	CurvePreferences: []tls.CurveID{tls.X25519, tls.CurveP256},
	//}
	return &Server{
		ServerHTTP: &http.Server{
			Addr:    ":" + cfg.Port,
			Handler: r,
			//TLSConfig: tlsConfig,
		},
	}
}

func (s *Server) Start(ctx context.Context) error {
	go func() {
		if err := s.ServerHTTP.ListenAndServe(); err != nil {
			log.Println(err)
			return
		}
	}()
	//"./certs/cert.pem", "./certs/key.pem"
	<-ctx.Done()
	return s.ServerHTTP.Shutdown(context.Background())
}

func (s *Server) Shutdown(ctx context.Context) error {
	return s.ServerHTTP.Shutdown(ctx)
}
