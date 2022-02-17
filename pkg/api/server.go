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

type ApiServer interface {
	Init(opts ...Options) error
	Start(ctx context.Context) error
}

type Server[H http.Handler, A ApiServer] struct {
	serv    *http.Server
	apiServ A
	opts    []Options
}

func NewServer[H http.Handler, A ApiServer](h H, apiServ A, addr string) (*Server[H, A], error) {
	srv := &http.Server{
		Addr:    addr,
		Handler: h,
	}
	opts := make([]Options, 0)
	opts = append(opts, WithAddr(addr))
	return &Server[H, A]{serv: srv, apiServ: apiServ, opts: opts}, nil
}

func (s *Server[T, A]) Start(ctx context.Context) error {
	if err := s.apiServ.Init(s.opts...); err != nil {
		return err
	}

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
