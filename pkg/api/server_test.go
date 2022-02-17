package api

import (
	"context"
	"testing"
	"time"

	"github.com/gin-gonic/gin"
)

var _ ApiServer = &mockApiServer{}

type mockApiServer struct {
	*Option
}

func (m mockApiServer) Init(opts ...Options) error {
	for _, f := range opts {
		f(m.Option)
	}
	return nil
}

func (m mockApiServer) Start(ctx context.Context) error {
	return nil
}

func Test_mock_server(t *testing.T) {
	route := gin.New()
	mas := &mockApiServer{}
	server, err := NewServer(route, mas, ":8080")
	if err != nil {
		t.Failed()
	}
	ctx, cancel := context.WithTimeout(context.Background(), 1)
	go server.Start(ctx)

	time.Sleep(time.Second)
	cancel()
}
