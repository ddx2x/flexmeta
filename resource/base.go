package resource

import (
	"github.com/laik/flexmeta/pkg/core"
)

type Base struct {
	core.Metadata `json:"metadata"`
	core.Spec     `json:"spec"`
}
