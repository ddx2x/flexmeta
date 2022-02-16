package store

import (
	"context"
	"errors"
)

var MockExpectListError = errors.New("MockExpectListError")

type MockStore[K comparable, Q map[K]any, O any] struct{}

func (m MockStore[K, Q, O]) Create(ctx context.Context, q Q, r O) (O, error) {
	return r, nil
}
func (m MockStore[K, Q, O]) List(ctx context.Context, q Q) ([]O, error) {
	return nil, MockExpectListError
}
func (m MockStore[K, Q, O]) Get(ctx context.Context, q Q) (O, error) {
	var ret O
	return ret, nil
}
func (m MockStore[K, Q, O]) Update(ctx context.Context, q Q, old O, new O, force bool) (O, error) {
	return new, nil
}
