package base

import (
	"context"

	"github.com/ddx2x/flexmeta/pkg/core"
	"github.com/gin-gonic/gin"
)

func (s *Server) pong(c *gin.Context) {
	c.JSON(200, gin.H{
		"message": "pong",
	})
}

func (s *Server) getAccount(g *gin.Context) {
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	targets, err := s.accont.List(ctx, Q{})
	if err != nil {
		g.JSON(500, map[string]interface{}{"error": err.Error()})
		return
	}
	objects := make([]core.Object[A], 0)
	for _, target := range targets {
		objects = append(objects, core.Object[A]{Item: target})
	}

	g.JSON(200, objects)
}
