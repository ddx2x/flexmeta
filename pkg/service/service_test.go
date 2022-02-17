package service

import (
	"context"
	"testing"

	"github.com/ddx2x/flexmeta/pkg/core"
	"github.com/ddx2x/flexmeta/pkg/store"
)

func Test_Base_Mock_Service2(t *testing.T) {
	type K string
	type Q map[K]any

	type E struct {
		core.Metadata `json:"metadata"`
		core.Spec     `json:"spec"`
	}
	type R = core.Object[E]

	// 定义一个BaseService拓展基础的Service
	type BaseService struct {
		Service[K, Q, E, R]
	}

	// 创建一core.Object[T]的实例
	base := core.Object[E]{}

	// 初始化一个BaseService
	baseService := &BaseService{}
	// 定义store为MockServiceStore
	baseService.Set(store.NewStore[K, Q, R](&MockServiceStore[K, Q, R]{}, nil))

	ctx := context.Background()

	// 创建
	err := baseService.Create(ctx, base)
	if err != MockExpectCreateError {
		t.Failed()
	}

	// 查询
	_, err = baseService.List(ctx, Q{"A": "B"})
	if err != MockExpectListError {
		t.Failed()
	}

}
