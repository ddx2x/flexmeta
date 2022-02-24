package base

import (
	"context"
	"fmt"
	"io"
	"time"

	"github.com/ddx2x/flexmeta/pkg/core"
	"github.com/ddx2x/flexmeta/pkg/log"
	"github.com/ddx2x/flexmeta/pkg/resource/account"
	"github.com/gin-gonic/gin"
)

type Boss struct {
	Uid     string `json:"uid"`
	Version string `json:"version"`
	Kind    string `json:"kind"`
	Vin     string `json:"vin"`
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

func (s *Server) _watch(ctx context.Context) <-chan error {
	return nil
}

func (s *Server) watch(c *gin.Context) {
	ctx, cancel := context.WithCancel(context.Background())
	_ = ctx
	ticker := time.NewTicker(time.Second * 2)
	index := 321

	l := log.G(ctx).WithFields(log.Fields{"watch": "process"})
	l.Infof("watch process start")

	c.Stream(func(w io.Writer) bool {
		select {
		case <-c.Writer.CloseNotify(): //client close
			cancel()
			l.Info("watch process close")
			return false
		case <-ticker.C:
			index++
			e := core.Event{}
			e.Type = core.ADDED
			e.Object = &Boss{
				Uid:     "123",
				Version: fmt.Sprintf("%d", index),
				Kind:    "boss",
				Vin:     fmt.Sprintf("粤A%d", index),
			}
			c.SSEvent("", e)
		}
		return true
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
