package base

import (
	"github.com/gin-gonic/gin"
)

func (server *Server) routes() {
	server.GET("/", server.Get)
}

func (s *Server) Get(g *gin.Context) {
	// var target *resource.Base
	// ctx, cancel := context.WithCancel(context.Background())
	// defer cancel()

	// target, err := s.base.Get(ctx, "a")
	// if err != nil {
	// 	g.JSON(500, err)
	// }

	// g.JSON(200, target)
}
