package service

import (
	"context"

	"github.com/laik/flexmeta/pkg/core"
	"github.com/laik/flexmeta/pkg/store"
)

type Service[T core.Objectable, S store.IStore[string, map[string]any, T]] struct {
	store S
}

func NewService[T core.Objectable, S store.IStore[string, map[string]any, T]](obj T, store S) *Service[T, S] {
	return &Service[T, S]{store}
}

func (s *Service[T, S]) Get(ctx context.Context, name string) (*T, error) {
	var target T
	target, err := s.store.Get(ctx, map[string]any{"name": name})
	if err != nil {
		return nil, err
	}
	return &target, nil
}
