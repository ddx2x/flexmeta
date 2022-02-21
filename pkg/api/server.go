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
	http.Handler
	Init(opts ...Options)
}

type Server struct {
	serv    *http.Server
	apiServ ApiServer
	opts    []Options
}

func NewServer(h ApiServer, addr string) *Server {
	srv := &http.Server{
		Addr:    addr,
		Handler: h,
	}
	opts := make([]Options, 0)
	opts = append(opts, WithAddr(addr))

	return &Server{
		serv:    srv,
		apiServ: h,
		opts:    opts,
	}
}

func (s *Server) Start(ctx context.Context) error {
	s.apiServ.Init(s.opts...)

	var errc = make(chan error)
	go func() {
		if err := s.serv.ListenAndServe(); err != nil {
			errc <- err
		}
	}()

	go func() {
		<-ctx.Done()
		s.serv.Shutdown(ctx)
		errc <- errors.New("server shutdown")
	}()

	return <-errc
}
