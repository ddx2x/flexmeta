package base

import (
	"github.com/gin-gonic/gin"
	"github.com/laik/flexmeta/pkg/api"
)

type BaseAPIServer struct {
	server api.Server[*gin.Engine]
}
