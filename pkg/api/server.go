package api

import (
	"context"
	"errors"
	"net/http"
)

type ApiServer interface {
	http.Handler
	Init(opts ...Options) error
}

type Server struct {
	serv    *http.Server
	apiServ ApiServer
	opts    []Options
}

func NewServer(h ApiServer, opts ...Options) *Server {
	opt := &Option{}
	for _, f := range opts {
		f(opt)
	}
	if opt.Addr == "" {
		opt.Addr = ":8080"
	}

	return &Server{
		serv:    &http.Server{Handler: h,Addr:  opt.Addr},
		apiServ: h,
		opts:    opts,
	}
}

func (s *Server) Start(ctx context.Context) error {
	if err := s.apiServ.Init(s.opts...); err != nil {
		return err
	}

	var errc = make(chan error)
	go func() {
		for {
			select {
			case <-ctx.Done():
				s.serv.Shutdown(ctx)
				errc <- errors.New("server shutdown")
				return
			default:
				if err := s.serv.ListenAndServe(); err != nil {
					errc <- err
				}
			}
		}
	}()

	return <-errc
}
