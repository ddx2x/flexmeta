package api

import (
	"context"
	"testing"
	"time"

	"github.com/gin-gonic/gin"
)

func Test_mock_server(t *testing.T) {
	route := gin.New()
	server, err := NewServer(route, ":8080")
	if err != nil {
		t.Failed()
	}
	ctx, cancel := context.WithTimeout(context.Background(), 1)
	go server.Start(ctx)

	time.Sleep(time.Second)
	cancel()
}
