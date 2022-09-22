package httpserver

import (
	"context"
	"net/http"
	"time"
)

const (
	defaultReadTimeout     = 5 * time.Second
	defaultWriteTimeout    = 5 * time.Second
	defaultAddr            = ":80"
	defaultShutdownTimeout = 3 * time.Second
)

type Server struct {
	server *http.Server
	notify chan error
	shutdownTimeout time.Duration
}

func NewServer(handler http.Handler, opts ...Option) *Server {
	httpServer := &http.Server{
		Addr: defaultAddr,
		Handler: handler,
		ReadTimeout: defaultReadTimeout,
		WriteTimeout: defaultWriteTimeout,
	}

	s := &Server{
		server: httpServer,
		notify: make(chan error, 1),
		shutdownTimeout: defaultShutdownTimeout,
	}

	for _, opt := range opts {
		opt(s)
	}

	return s
}

func (s *Server) Start() {
	go func() {
		s.notify <- s.server.ListenAndServe()
		close(s.notify)
	}()
}

func (s *Server) Notify() <-chan error {
	return s.notify
}

func (s *Server) Shutdown() error {
	ctx, cancel := context.WithTimeout(context.Background(), s.shutdownTimeout)
	defer cancel()

	return s.server.Shutdown(ctx)
}
