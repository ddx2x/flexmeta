package main

import (
	"context"

	"github.com/ddx2x/flexmeta/pkg/api"
	"github.com/ddx2x/flexmeta/pkg/api/base"
	"github.com/ddx2x/flexmeta/pkg/log"
	"github.com/ddx2x/flexmeta/pkg/logrus"
	"github.com/ddx2x/flexmeta/pkg/signals"
	"github.com/sirupsen/logrus"
)

func main() {
	stopCh := signals.SetupSignalHandler()
	ctx, cancel := context.WithCancel(context.Background())
	go func() {
		<-stopCh
		cancel()
	}()

	log.L = logrus.FromLogrus(logrus.NewEntry(logrus.StandardLogger()))
	log.G(ctx).Info("start cr webserver")

	if err := api.NewServer(&base.Server{}, ":8080").Start(ctx); err != nil {
		panic(err)
	}
}
