package service

import "github.com/laik/flexmeta/pkg/core"

type MockServiceObject struct {
	core.Metadata `json:"metadata"`
	core.Spec     `json:"spec"`
}
