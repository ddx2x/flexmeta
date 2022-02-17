package base

import (
	"context"

	"github.com/gin-gonic/gin"
	"github.com/laik/flexmeta/pkg/api"
	"github.com/laik/flexmeta/pkg/core"
	"github.com/laik/flexmeta/pkg/service"
	"github.com/laik/flexmeta/pkg/store"
	rc "github.com/laik/flexmeta/resource"
)

type (
	K  string
	Q  map[K]any
	B  rc.Base
	RB = core.Object[B]
)

type Server struct {
	*gin.Engine
	base service.Service[K, Q, B, RB]
}

func (s *Server) Init(opts ...api.Options) error {
	engine := gin.Default()
	option := &api.Option{}
	for _, f := range opts {
		f(option)
	}
	server := &Server{
		Engine: engine,
	}

	server.base.Set(store.NewStore[K, Q, RB](&MockBaseStore[K, Q, RB]{}, nil))

	server.routes()
	
	return nil
}

func (s *Server) Start(ctx context.Context) error {
	return nil
}
