package base

import (
	"context"

	"github.com/ddx2x/flexmeta/pkg/core"
	"github.com/ddx2x/flexmeta/pkg/resource/account"
	"github.com/gin-gonic/gin"
)

type Boss struct {
	Uid     string `json:"uid"`
	Version string `json:"version"`
	Kind    string `json:"kind"`
}

func (s *Server) welcome(c *gin.Context) {
	c.JSON(200, []*Boss{
		{
			Uid:     "123",
			Version: "321",
			Kind:    "boss",
		}})
}

func (s *Server) pong(c *gin.Context) {
	c.JSON(200, gin.H{
		"message": "pong",
	})
}

func (s *Server) getAccount(g *gin.Context) {
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	targets, err := s.a.List(ctx, Q{})
	if err != nil {
		g.JSON(500, map[string]interface{}{"error": err.Error()})
		return
	}
	objects := make([]core.Object[account.Account], 0)
	for _, target := range targets {
		objects = append(objects, core.Object[account.Account]{Item: target})
	}

	g.JSON(200, objects)
}
