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
)

type Server struct {
	*gin.Engine
	opt *api.Option

	b *service.Service[K, Q, base.Base]
	a *service.Service[K, Q, account.Account]
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
	setKVStore(ctx, option.StoreAddr, s.b)
	// init account service
	setKVStore(ctx, option.StoreAddr, s.a)

	//route register
	{
		s.GET("/", s.pong)
		s.GET("/account", s.getAccount)
	}

	return nil
}

func setKVStore[K comparable, Q ~map[K]any, R any](ctx context.Context, addr string, s *service.Service[K, Q, R]) error {
	m, err := mongo.NewMongoCli[K, Q, R](ctx, addr)
	if err != nil {
		return err
	}
	s.Set(store.NewStore[K, Q, R](m, nil))
	return nil
}
