package base

import (
	"context"

	"github.com/ddx2x/flexmeta/pkg/api"
	"github.com/ddx2x/flexmeta/pkg/resource/account"
	"github.com/ddx2x/flexmeta/pkg/resource/base"
	"github.com/ddx2x/flexmeta/pkg/service"
	"github.com/ddx2x/flexmeta/pkg/store"
	"github.com/ddx2x/flexmeta/pkg/store/mongo"
	"github.com/gin-gonic/gin"
)

type (
	K string
	Q map[K]any
	B base.Base
	A account.Account
)

type Server struct {
	*gin.Engine
	opt *api.Option

	base   service.Service[K, Q, B, B]
	accont service.Service[K, Q, A, A]
}

func NewServer() *Server {
	return &Server{
		Engine: gin.New(),
		opt:    &api.Option{},
	}
}

func (s *Server) Init(opts ...api.Options) error {
	option := &api.Option{}
	for _, f := range opts {
		f(option)
	}

	ctx := context.Background()

	// init base service
	m, err := mongo.NewMongoCli[K, Q, B](ctx, option.StoreAddr)
	if err != nil {
		return err
	}

	s.base.Set(store.NewStore[K, Q, B](m, nil))

	// init account service
	m1, err := mongo.NewMongoCli[K, Q, A](ctx, option.StoreAddr)
	if err != nil {
		return err
	}
	s.accont.Set(store.NewStore[K, Q, A](m1, nil))

	//route register
	{
		s.GET("/", s.pong)

		s.GET("/account", s.getAccount)
	}

	return nil
}
