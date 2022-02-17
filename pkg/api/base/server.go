package base

import (
	"context"

	"github.com/ddx2x/flexmeta/pkg/api"
	"github.com/ddx2x/flexmeta/pkg/core"
	"github.com/ddx2x/flexmeta/pkg/service"
	"github.com/ddx2x/flexmeta/pkg/store"
	rc "github.com/ddx2x/flexmeta/resource"
	"github.com/gin-gonic/gin"
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
