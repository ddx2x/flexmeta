package store

import (
	"context"
	"testing"
)



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
