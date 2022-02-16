package base

import (
	"context"

	"github.com/gin-gonic/gin"
	"github.com/laik/flexmeta/pkg/api"
	"github.com/laik/flexmeta/pkg/service"
	"github.com/laik/flexmeta/pkg/store"
	"github.com/laik/flexmeta/resource"
)

type Server struct {
	serv   *api.Server[*gin.Engine]
	engine *gin.Engine
	base   *service.Service[resource.Base, store.IStore[string, map[string]any, resource.Base]]

	registers []func()
}

func RunServer(ctx context.Context, addr string) error {
	engine := gin.Default()

	apiServer, err := api.NewServer(engine, addr)
	if err != nil {
		return err
	}

	server := &Server{
		engine:    engine,
		serv:      apiServer,
		registers: make([]func(), 0),
	}

	server.routes()

	return server.serv.Start(ctx)
}
