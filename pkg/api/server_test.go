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
	*gin.Engine
}

func (m mockApiServer) Init(opts ...Options) {
	for _, f := range opts {
		f(m.Option)
	}
}

func Test_mock_server(t *testing.T) {
	mas := &mockApiServer{Engine: gin.New()}
	server := NewServer(mas, ":8080")
	
	ctx, cancel := context.WithTimeout(context.Background(), 1)
	go server.Start(ctx)

	time.Sleep(time.Second)
	cancel()
}
