package base

import (
	"context"

	"github.com/gin-gonic/gin"
)

func (server *Server) routes() {
	server.GET("/", server.Get)
}

func (s *Server) Get(g *gin.Context) {
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	target, err := s.base.Get(ctx, Q{"name": "a"})
	if err != nil {
		g.JSON(500, err)
	}

	g.JSON(200, target)
}
