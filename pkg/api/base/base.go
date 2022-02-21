package base

import (
	"context"

	"github.com/ddx2x/flexmeta/pkg/core"
	"github.com/gin-gonic/gin"
)

func (server *Server) routes() {
	server.GET("/", server.Get)
	server.GET("/ping", server.Get)
}

func (s *Server) pong(c *gin.Context) {
	c.JSON(200, gin.H{
		"message": "pong",
	})
}

func (s *Server) Get(g *gin.Context) {
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	target, err := s.base.List(ctx, Q{"name": "a"})
	if err != nil {
		g.JSON(500, err)
	}

	g.JSON(200, (&core.Items[B]{}).From(target))
}
