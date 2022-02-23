package base

import (
	"github.com/ddx2x/flexmeta/pkg/core"
)

type Spec struct {
	UserId string `json:"userId"`
}

type Base struct {
	core.Metadata `json:"metadata"`
	Spec     `json:"spec"`
}