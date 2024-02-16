package http

import (
	"context"
	"errors"
	"net"
	"net/http"

	"github.com/asawo/api/logger"
	"github.com/asawo/api/service"
)

// Server is a http server that expose ws endpoint as well as the hooks
type Server struct {
	mux     *http.ServeMux
	server  *http.Server
	logger  *logger.Log
	service service.Service
}

// New creates a new http server
func New(ctx context.Context, logger *logger.Log, service service.Service) (*Server, error) {
	server := &Server{
		mux:     http.NewServeMux(),
		logger:  logger,
		service: service,
	}

	server.registerHandlers(ctx)
	return server, nil
}

// Serve starts the http server given listener
func (s *Server) Serve(ln net.Listener) error {
	server := &http.Server{
		Handler: s.mux,
	}
	s.server = server

	if err := server.Serve(ln); err != nil && !errors.Is(err, http.ErrServerClosed) {
		return err
	}

	return nil
}

// GracefulStop tries to stop server process without affecting existing requests
func (s *Server) GracefulStop(ctx context.Context) error {
	return s.server.Shutdown(ctx)
}

// registerHandlers registers http endpoints and their handlers
func (s *Server) registerHandlers(ctx context.Context) {
	s.mux.Handle("/api/invoices", s.invoicesHandler(ctx))
}
