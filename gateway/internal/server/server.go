package server

import (
	"fmt"
	"github.com/ArtyomYatsenko/gateway/internal/config"
	"go.uber.org/zap"
	"net/http"
)

type Server struct {
	httpServer *http.Server
	logger     *zap.Logger
}

func (s *Server) Start(config config.Server, handler http.Handler) error {
	s.httpServer = &http.Server{
		Addr:         config.Address + ":" + config.Port,
		Handler:      handler,
		WriteTimeout: config.WriteTimeout,
		ReadTimeout:  config.ReadTimeout,
	}

	fmt.Printf("%+v", s.httpServer)
	return s.httpServer.ListenAndServe()
}

func (s *Server) Stop() error {
	return nil
}
