package base

import (
	"context"
	"testing"

	"github.com/laik/flexmeta/pkg/store"
)

func Test_Base_Service(t *testing.T) {
	s := &store.MockStore[string, map[string]any, Base]{}
	store, err := store.NewStore[string, map[string]any, Base](*s)
	if err != nil {
		t.Fatalf("%s", err)
	}
	if store == nil {
		t.Failed()
	}
	ctx := context.Background()
	_, err = store.List(ctx, map[string]any{"a": "b"})

	t.Logf("ok")
}
