package apiserver

import (
	"EmployeeServiceWithQuickSortXml/config"
	"EmployeeServiceWithQuickSortXml/internal/handler"
	"context"
	"fmt"
	"log"
	"net/http"
	"time"
)

type Server struct {
	httpServer *http.Server
}

func New(config *config.Config, handler *handler.Handler) *Server {
	return &Server{
		httpServer: &http.Server{
			Addr:    fmt.Sprintf(":%d", config.Port),
			Handler: handler.Router,
			//BaseContext: func(_ net.Listener) context.Context {
			//	return ctx
			//},
			MaxHeaderBytes: 1 << 20,
			ReadTimeout:    10 * time.Second,
			WriteTimeout:   10 * time.Second,
		},
	}
}

func (s *Server) Run() {
	if err := s.httpServer.ListenAndServe(); err != nil {
		log.Fatalf("%v", err)
		return

	}
}

func (s *Server) Shutdown(ctx context.Context) error {
	return s.httpServer.Shutdown(ctx)
}
