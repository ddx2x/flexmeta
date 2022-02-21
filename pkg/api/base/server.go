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
	OB = core.Object[B]
)

type Server struct {
	*gin.Engine
	opt *api.Option

	base service.Service[K, Q, B, OB]
}

func NewServer() *Server {
	return &Server{
		Engine: gin.New(),
		opt:    &api.Option{},
	}
}

func (s *Server) Init(opts ...api.Options) {
	option := &api.Option{}
	for _, f := range opts {
		f(option)
	}

	s.base.Set(store.NewStore[K, Q, OB](&MockBaseStore[K, Q, OB]{}, nil))

	{
		s.GET("/", s.pong)
	}
}
