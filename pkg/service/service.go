package service

import (
	"github.com/ddx2x/flexmeta/pkg/core"
	"github.com/ddx2x/flexmeta/pkg/store"
)

type Serviceable[K comparable, Q ~map[K]any, E core.Objectizable, R core.Object[E]] interface {
	~struct {
		Service[K, Q, E, R]
	}
}
type Service[K comparable, Q ~map[K]any, E core.Objectizable, R core.Object[E]] struct {
	*store.Store[K, Q, R]
}

func (s *Service[K, Q, E, R]) Set(store *store.Store[K, Q, R]) {
	s.Store = store
}
