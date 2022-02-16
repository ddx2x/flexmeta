package store

import (
	"context"
	"errors"
	"testing"
)

var (
	MockExpectListError = errors.New("MockExpectListError")
)

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

func Test_Store_NewStore(t *testing.T) {
	s := &MockStore[string, map[string]any, interface{}]{}
	store, err := NewStore[string, map[string]any, interface{}](*s)
	if err != nil {
		t.Fatalf("%s", err)
	}
	if store == nil {
		t.Failed()
	}
	t.Logf("ok")

	ctx := context.Background()
	_, err = store.List(ctx, map[string]any{"a": "b"})
	if err != MockExpectListError {
		t.Failed()
	}
}

func Test_Store_Example1(t *testing.T) {
	s := &MockStore[string, map[string]any, any]{}
	store, err := NewStore[string, map[string]any, any](*s)
	if err != nil {
		t.Fatalf("%s", err)
	}
	if store == nil {
		t.Failed()
	}
	t.Logf("ok")

	ctx := context.Background()
	_, err = store.List(ctx, map[string]any{"a": "b"})
	if err != MockExpectListError {
		t.Failed()
	}
}
