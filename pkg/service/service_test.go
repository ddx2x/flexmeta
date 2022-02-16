package service

import (
	"context"
	"testing"

	"github.com/laik/flexmeta/pkg/store"
)

func Test_Base_Mock_Service2(t *testing.T) {
	store := &store.MockStore[string, map[string]any, MockServiceObject]{}
	service := NewService(MockServiceObject{}, store)
	service.Get(context.Background(), "a")
}
