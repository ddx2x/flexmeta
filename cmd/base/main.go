package main

import (
	"context"

	"github.com/laik/flexmeta/pkg/api/base"
	"github.com/laik/flexmeta/pkg/signals"
)

func main() {
	stopCh := signals.SetupSignalHandler()
	ctx, cancel := context.WithCancel(context.Background())
	go func() {
		<-stopCh
		cancel()
	}()

	if err := base.RunServer(ctx, ":8080"); err != nil {
		panic(err)
	}
}
