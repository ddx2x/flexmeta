package base

import (
	"github.com/ddx2x/flexmeta/pkg/api"
	"github.com/ddx2x/flexmeta/pkg/core"
	rc "github.com/ddx2x/flexmeta/pkg/resource"
	"github.com/ddx2x/flexmeta/pkg/service"
	"github.com/ddx2x/flexmeta/pkg/store"
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

func (s *Server) Init(opts ...api.Options) {
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
}
