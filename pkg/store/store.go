package store

import (
	"context"
	"fmt"

	"github.com/ddx2x/flexmeta/pkg/core"
)

var DataNotFound = fmt.Errorf("dataNotFound")

// IStore interface method for internal implementation, only KV storage or rdbms storage is implemented here
type IKVStore[K comparable, Q ~map[K]any, R any] interface {
	// Create Object
	Create(context.Context, R) error
	//
	Update(ctx context.Context, old R, new R, q Q) (R, error)
	//
	Delete(context.Context, Q) error
	//
	List(context.Context, Q) ([]R, error)
	//
	Get(context.Context, Q) (R, error)
	//
	Watch(context.Context, Q) (<-chan core.Event, <-chan error)
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

func Get[K comparable, Q ~map[K]any, R any](store interface{}) *Store[K, Q, R] {
	return store.(*Store[K, Q, R])
}
