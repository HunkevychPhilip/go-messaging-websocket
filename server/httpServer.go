package server

import (
	"context"
	"net/http"
	"time"
)

type HTTPServer struct {
	httpServer *http.Server
}

func NewHTTPServer(port string, h http.Handler) *HTTPServer {
	return &HTTPServer{
		&http.Server{
			Addr:           ":" + port,
			Handler:        h,
			ReadTimeout:    10 * time.Second,
			WriteTimeout:   10 * time.Second,
			MaxHeaderBytes: 1 << 20,
		},
	}
}

func (s *HTTPServer) Start() error {
	return s.httpServer.ListenAndServe()
}

func (s *HTTPServer) ShutDown(ctx context.Context) error {
	return s.httpServer.Shutdown(ctx)
}
