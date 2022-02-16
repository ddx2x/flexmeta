package base

import (
	"context"

	"github.com/laik/flexmeta/pkg/core"
	"github.com/laik/flexmeta/pkg/service"
	"github.com/laik/flexmeta/pkg/store"
)

type Base struct {
	core.Metadata `json:"metadata"`
	core.Spec     `json:"spec"`
}

type BaseService struct {
	service *service.Service[Base, store.MockStore[string, map[string]any, Base]]
}

func (s *BaseService) Get(ctx context.Context, name string) (*Base, error) {
	return s.service.Get(ctx, name)
}
