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

	// log.L = logadapter.FromLogrus(logrus.NewEntry(logrus.StandardLogger()))
	// log.G(ctx).Info("start cr webserver")

	if err := api.NewServer(base.NewServer(), ":8080").Start(ctx); err != nil {
		panic(err)
	}
}
