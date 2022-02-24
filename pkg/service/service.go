package service

import (
	"github.com/ddx2x/flexmeta/pkg/core"
	"github.com/ddx2x/flexmeta/pkg/store"
)

// type Serviceable[K comparable, Q ~map[K]any, R core.IObject] interface {
// 	~struct {
// 		Service[K, Q, R]
// 	}
// }

type Watcher interface {
}

type Service[K comparable, Q ~map[K]any, R core.IObject] struct {
	*store.Store[K, Q, R]
}

func (s *Service[K, Q, R]) Watchs(watchers ...Watcher) error {

	return nil
}

func (s *Service[K, Q, R]) Set(store *store.Store[K, Q, R]) {
	s.Store = store
}
