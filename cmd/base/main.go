package main

import (
	"context"

	"github.com/ddx2x/flexmeta/pkg/api"
	"github.com/ddx2x/flexmeta/pkg/api/base"
	"github.com/ddx2x/flexmeta/pkg/signals"
)

func main() {
	stopCh := signals.SetupSignalHandler()
	ctx, cancel := context.WithCancel(context.Background())
	go func() {
		<-stopCh
		cancel()
	}()

	bs := &base.Server{}
	apiServer, err := api.NewServer(bs, bs, ":8080")
	if err != nil {
		panic(err)
	}

	if err = apiServer.Start(ctx); err != nil {
		panic(err)
	}
}
