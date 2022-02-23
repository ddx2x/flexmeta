package service

import (
	"github.com/ddx2x/flexmeta/pkg/store"
)

type Serviceable[K comparable, Q ~map[K]any, R any] interface {
	~struct {
		Service[K, Q, R]
	}
}
type Service[K comparable, Q ~map[K]any, R any] struct {
	*store.Store[K, Q, R]
}

func (s *Service[K, Q, R]) Set(store *store.Store[K, Q, R]) {
	s.Store = store
}
