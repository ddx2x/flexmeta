package base

import (
	"context"
	"fmt"
	"io"
	"math/rand"
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

type Options struct {
	XAxis []string `json:"xAxis"`
	YAxis []string `json:"yAxis"`
	Data  [][]int  `json:"data"`
}
type View struct {
	Uid     string `json:"uid"`
	Version string `json:"version"`
	Kind    string `json:"kind"`
	Options `json:"options"`
}

func (s *Server) welcome(c *gin.Context) {
	c.JSON(200, []*Boss{
		{
			Uid:     "123",
			Version: "321",
			Kind:    "boss",
		}})
}

func iter(i int) []struct{} {
	return make([]struct{}, i)
}

var random = func() int { return int(rand.NewSource(time.Now().UnixNano()).Int63()) }

func (s *Server) view(c *gin.Context) {
	views := make([]*View, 0)
	for i := range iter(100) {
		data := [][]int{
			{random(), random(), random(), random()},
			{random(), random(), random(), random()},
			{random(), random(), random(), random()},
			{random(), random(), random(), random()},
		}
		views = append(views, &View{
			Uid:     fmt.Sprintf("%d", i),
			Version: "1",
			Kind:    "view",
			Options: Options{
				XAxis: []string{"2017-10-24", "2017-10-25", "2017-10-26", "2017-10-27"},
				Data:  data,
			},
		})
	}
	c.JSON(200, views)
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
	index := 1

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

			for i := range iter(100) {
				data := [][]int{
					{random(), random(), random(), random()},
					{random(), random(), random(), random()},
					{random(), random(), random(), random()},
					{random(), random(), random(), random()},
				}
				view := &View{
					Uid:     fmt.Sprintf("%d", i),
					Version: fmt.Sprintf("%d", index),
					Kind:    "view",
					Options: Options{
						XAxis: []string{"2017-10-24", "2017-10-25", "2017-10-26", "2017-10-27"},
						Data:  data,
					},
				}
				e.Object = view
				c.SSEvent("", e)
			}
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
