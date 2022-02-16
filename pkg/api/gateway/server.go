package gateway

import (
	"github.com/gin-gonic/gin"
	"github.com/laik/flexmeta/pkg/api"
)

type Server struct {
	serv api.Server[*gin.Engine]
}
