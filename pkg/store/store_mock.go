package store

import (
	"context"
	"errors"
)

var MockExpectListError = errors.New("MockExpectListError")

type MockStore[K comparable, Q ~map[K]any, R any] struct{}

// Create Object
func (m MockStore[K, Q, R]) Create(context.Context, R) error { return nil }

//
func (m MockStore[K, Q, R]) Update(ctx context.Context, old R, new R, force bool) (R, error) {
	return new, nil
}
//
func (m MockStore[K, Q, R]) Delete(context.Context, Q) error { return nil }

//
func (m MockStore[K, Q, R]) List(context.Context, Q) ([]R, error) { 
	return nil, MockExpectListError 
}

//
func (m MockStore[K, Q, R]) Get(context.Context, Q) (R, error) {
	var r R
	return r, nil
}
