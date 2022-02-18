package store

import (
	"context"
)

// IStore interface method for internal implementation, only KV storage or rdbms storage is implemented here

type IKVStore[K comparable, Q ~map[K]any, R any] interface {
	// Create Object
	Create(context.Context, R) error
	//
	Update(ctx context.Context, old R, new R, force bool) (R, error)
	//
	Delete(context.Context, Q) error
	//
	List(context.Context, Q) ([]R, error)
	//
	Get(context.Context, Q) (R, error)
	//
	// Watch(context.Context, Q) (<-chan R,<-chan error)
}

type IRDBMSStore[K comparable, Q ~map[K]any, R any] interface {
}

type Store[K comparable, Q ~map[K]any, R any] struct {
	IKVStore[K, Q, R]
	IRDBMSStore[K, Q, R]
}

func NewStore[K comparable, Q ~map[K]any, R any](kv IKVStore[K, Q, R], rdbms IRDBMSStore[K, Q, R]) *Store[K, Q, R] {
	return &Store[K, Q, R]{kv, rdbms}
}
