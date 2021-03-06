package service

import (
	"context"
	"errors"

	"github.com/ddx2x/flexmeta/pkg/core"
)

var (
	MockExpectCreateError = errors.New("MockExpectCreateError")
	MockExpectListError   = errors.New("MockExpectListError")
	MockExpectUpdateError = errors.New("MockExpectCreateError")
	MockExpectDeleteError = errors.New("MockExpectDeleteError")
)

type MockServiceStore[K comparable, Q ~map[K]any, R core.IObject] struct{}

func (m MockServiceStore[K, Q, R]) Create(context.Context, R) error {
	return MockExpectCreateError
}

func (m MockServiceStore[K, Q, R]) Update(ctx context.Context, old R, new R, q Q) (R, error) {
	return new, MockExpectUpdateError
}

func (m MockServiceStore[K, Q, R]) Delete(context.Context, Q) error {
	return MockExpectDeleteError
}

func (m MockServiceStore[K, Q, R]) List(context.Context, Q) ([]R, error) {
	return nil, MockExpectListError
}

func (m MockServiceStore[K, Q, R]) Get(context.Context, Q) (R, error) {
	var r R
	return r, nil
}

func (m MockServiceStore[K, Q, R]) Watch(context.Context, Q) (<-chan core.Event, <-chan error) {
	return nil, nil
}
