package resource

import (
	"github.com/ddx2x/flexmeta/pkg/core"
)

type Base struct {
	core.Metadata `json:"metadata"`
	core.Spec     `json:"spec"`
}
