package api

import (
	"context"
	"errors"
	"net/http"
)

type HttpServer interface {
	ServeHTTP(w http.ResponseWriter, r *http.Request)
	Run(addr ...string) error
}
type Server[H http.Handler] struct {
	serv *http.Server
}

func NewServer[H http.Handler](handle H, addr string) (*Server[H], error) {
	srv := &http.Server{
		Addr:    addr,
		Handler: handle,
	}
	return &Server[H]{srv}, nil
}

func (s *Server[T]) Start(ctx context.Context) error {
	var errc = make(chan error)
	go func() {
		if err := s.serv.ListenAndServe(); err != nil {
			errc <- err
		}
	}()
	go func() {
		for range ctx.Done() {
			errc <- errors.New("server shutdown")
		}
	}()
	return <-errc
}
