package main

import (
	"context"

	"github.com/laik/flexmeta/pkg/api"
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

	// mci, err := mongo.NewMongoCli[string, map[string]any, core.IObject](ctx, "mongodb://localhost:27017")
	// if err != nil {
	// 	panic(err)
	// }

	// store, err := store.NewStore[string, map[string]any, core.IObject](mci)
	// if err != nil {
	// 	panic(err)
	// }
	
	baseServ := &base.Server{}
	apiServer, err := api.NewServer(baseServ, baseServ, ":8080")
	if err != nil {
		panic(err)
	}

	if err = apiServer.Start(ctx); err != nil {
		panic(err)
	}
}
