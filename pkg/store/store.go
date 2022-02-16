package store

import (
	"context"
)

type IStore[K comparable, Q ~map[K]any, O any] interface {
	Create(ctx context.Context, q Q, o O) (O, error)
	Update(ctx context.Context, q Q, old O, new O, force bool) (O, error)
	List(ctx context.Context, q Q) ([]O, error)
	Get(ctx context.Context, q Q) (O, error)
}

type Store[K comparable, Q ~map[K]any, O any, S IStore[K, Q, O]] struct {
	store S
}

func NewStore[K comparable, Q ~map[K]any, O any, S IStore[K, Q, O]](s S) (*Store[K, Q, O, S], error) {
	return &Store[K, Q, O, S]{s}, nil
}

func (s *Store[K, Q, O, S]) List(ctx context.Context, q Q) ([]O, error) {
	//do something
	return s.store.List(ctx, q)
}
