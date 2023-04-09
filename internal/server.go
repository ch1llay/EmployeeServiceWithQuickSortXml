package apiserver

import (
	"EmployeeServiceWithQuickSortXml/config"
	"EmployeeServiceWithQuickSortXml/internal/handler"
	"context"
	"net/http"
	"time"
)

type Server struct {
	httpServer *http.Server
}

func New() *Server {
	return &Server{}
}

func (s *Server) Configure(config *config.Config, handler *handler.Handler) {
	s.httpServer = &http.Server{
		Addr:           ":" + string(config.Port),
		Handler:        handler.Router,
		MaxHeaderBytes: 1 << 20,
		ReadTimeout:    10 * time.Second,
		WriteTimeout:   10 * time.Second,
	}
}

func (s *Server) Shutdown(ctx context.Context) error {
	return s.httpServer.Shutdown(ctx)
}
