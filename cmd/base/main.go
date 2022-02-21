package main

import (
	"context"

	"github.com/ddx2x/flexmeta/pkg/api"
	"github.com/ddx2x/flexmeta/pkg/api/base"

	"github.com/ddx2x/flexmeta/pkg/signals"
)

type (
	K string
	Q map[K]interface{}
	R any
)

func main() {
	stopCh := signals.SetupSignalHandler()
	ctx, cancel := context.WithCancel(context.Background())

	go func() {
		<-stopCh
		cancel()
	}()
	
	if err := api.NewServer(
		base.NewServer(),
		api.Addr(":8080"),
		api.StoreAddr("mongodb://127.0.0.1:27017/admin"),
	).
		Start(ctx); err != nil {
		panic(err)
	}
}
