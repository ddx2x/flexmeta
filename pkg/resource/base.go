package resource

import (
	"github.com/ddx2x/flexmeta/pkg/core"
)

type BaseSpec struct {
	UserId string `json:"userId"`
}
type Base struct {
	core.Metadata `json:"metadata"`
	core.Spec     `json:"spec"`
}
