package main

import (
	"context"

	"github.com/ddx2x/flexmeta/pkg/api"
	"github.com/ddx2x/flexmeta/pkg/api/base"
	"github.com/ddx2x/flexmeta/pkg/log"
	rs "github.com/ddx2x/flexmeta/pkg/log/logrus"
	"github.com/ddx2x/flexmeta/pkg/signals"
	"github.com/sirupsen/logrus"
)

type (
	K string
	Q map[K]interface{}
	R any
)

func main() {
	stopCh := signals.SetupSignalHandler()
	ctx, cancel := context.WithCancel(context.Background())

	log.L = rs.FromLogrus(logrus.NewEntry(logrus.StandardLogger()))

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
