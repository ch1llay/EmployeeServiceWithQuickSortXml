package apiserver

import (
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

func (s *Server) Run(config *Config, handler *handler.Handler) error {
	s.httpServer = &http.Server{
		Addr:           config.BindAddr,
		Handler:        handler.Router,
		MaxHeaderBytes: 1 << 20,
		ReadTimeout:    10 * time.Second,
		WriteTimeout:   10 * time.Second,
	}

	if err := handler.EmployeeService.EmployeeRepository.InitDatabase(); err != nil {
		return err
	}

	return s.httpServer.ListenAndServe()
}

func (s *Server) Shutdown(ctx context.Context) error {
	return s.httpServer.Shutdown(ctx)
}
