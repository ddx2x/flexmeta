package resource

import (
	"context"
	"testing"

	"github.com/laik/flexmeta/pkg/service"
	"github.com/laik/flexmeta/pkg/store"
)

func Test_Base_Mock_Service2(t *testing.T) {
	store := &store.MockStore[string, map[string]any, Base]{}
	service := service.NewService(Base{}, store)
	service.Get(context.Background(), "a")
}
