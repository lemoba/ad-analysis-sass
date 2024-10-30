package http

import (
	"context"
	"errors"
	"fmt"
	"net/http"
	"time"

	"github.com/lemoba/ad-analysis-sass/pkg/log"
)

type Server struct {
	httpSrv *http.Server
	host    string
	port    int
	logger  *log.Logger
}

type Option func(*Server)

func NewServer(logger *log.Logger, opts ...Option) *Server {
	s := &Server{
		logger: logger,
	}

	for _, opt := range opts {
		opt(s)
	}

	return s
}

func WithServerHost(host string) Option {
	return func(s *Server) {
		s.host = host
	}
}

func WithServerPort(port int) Option {
	return func(s *Server) {
		s.port = port
	}
}

func (s *Server) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	// Your handler logic here
	w.WriteHeader(http.StatusOK)
	w.Write([]byte("Hello from Server"))
}

func (s *Server) Start(ctx context.Context) error {
	s.httpSrv = &http.Server{
		Addr:    fmt.Sprintf("%s:%d", s.host, s.port),
		Handler: s,
	}

	if err := s.httpSrv.ListenAndServe(); err != nil && !errors.Is(err, http.ErrServerClosed) {
		s.logger.Sugar().Fatalf("listen http server: %s\n", err)
	}

	return nil
}

func (s *Server) Stop(ctx context.Context) error {
	s.logger.Sugar().Info("Shutting down server...")

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	if err := s.httpSrv.Shutdown(ctx); err != nil {
		s.logger.Sugar().Fatalf("Server forced to shutdown: %s\n", err)
	}

	s.logger.Sugar().Info("Server shutdown ...")

	return nil
}
