package store

import (
	"context"
	"testing"
)

func Test_Store_NewStore(t *testing.T) {
	type K string
	type Q map[K]interface{}
	type R interface{}

	store := NewStore[K, Q, R](&MockStore[K, Q, R]{}, nil)
	if &store == nil {
		t.Failed()
	}

	ctx := context.Background()
	_, err := store.List(ctx, Q{"a": "b"})
	if err != MockExpectListError {
		t.Failed()
	}

	t.Logf("ok")
}
