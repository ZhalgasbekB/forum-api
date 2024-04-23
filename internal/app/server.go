package app

import (
	"context"
	"crypto/tls"
	"log"
	"net/http"

	"gitea.com/lzhuk/forum/internal/server"
	"gitea.com/lzhuk/forum/pkg/config"
)

type Server struct {
	ServerHTTP *http.Server
}

func NewServer(cfg config.Config, r server.Router) *Server {
	//certManager := autocert.Manager{
	//	Prompt: autocert.AcceptTOS,
	//	Cache:  autocert.DirCache("certs1"),
	//}
	//tlsConfig := certManager.TLSConfig()
	//tlsConfig.MinVersion = tls.VersionTLS12
	//tlsConfig.PreferServerCipherSuites = true
	//tlsConfig.CipherSuites = []uint16{
	//	tls.TLS_ECDHE_RSA_WITH_AES_256_GCM_SHA384,
	//	tls.TLS_ECDHE_RSA_WITH_AES_128_GCM_SHA256,
	//}
	//tlsConfig1 := &tls.Config{
	//	MinVersion:               tls.VersionTLS12,
	//	PreferServerCipherSuites: true,
	//	CipherSuites: []uint16{
	//		tls.TLS_ECDHE_RSA_WITH_AES_256_GCM_SHA384,
	//		tls.TLS_ECDHE_RSA_WITH_AES_128_GCM_SHA256,
	//		// Add more secure cipher suites here
	//	},
	//}
	//go http.ListenAndServe(":http", certManager.HTTPHandler(nil))
	//go http.ListenAndServe(":80", certManager.HTTPHandler(r))

	tlsConfig := &tls.Config{
		CurvePreferences: []tls.CurveID{tls.X25519, tls.CurveP256},
	}
	return &Server{
		ServerHTTP: &http.Server{
			Addr:      ":" + cfg.Port,
			Handler:   r,
			TLSConfig: tlsConfig,
		},
	}
}

func (s *Server) Start(ctx context.Context) error {
	go func() {
		if err := s.ServerHTTP.ListenAndServeTLS("./certs/cert.pem", "./certs/key.pem"); err != nil {
			log.Println(err)
			return
		}
		//"./certs/cert.pem", "./certs/key.pem"
		//log.Fatal(server.ListenAndServeTLS("", ""))
	}()
	<-ctx.Done()
	return s.ServerHTTP.Shutdown(context.Background())
}

func (s *Server) Shutdown(ctx context.Context) error {
	return s.ServerHTTP.Shutdown(ctx)
}
