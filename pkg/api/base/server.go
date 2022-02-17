package base

import (
	"context"

	"github.com/gin-gonic/gin"
	"github.com/laik/flexmeta/pkg/api"
)

type Server struct {
	*gin.Engine
}

func (s *Server) Init(opts ...api.Options) error {
	engine := gin.Default()
	option := &api.Option{}
	for _, f := range opts {
		f(option)
	}

	// base := service.NewService(resource.Base{}, nil)
	server := &Server{
		Engine: engine,
	}
	server.routes()
	return nil
}

func (s *Server) Start(ctx context.Context) error {
	return nil
}
