package core

import (
	"golang.org/x/exp/maps"
	"testing"
)

type Action struct {
	Metadata `json:"metadata"`
	Spec     `json:"spec"`
}

func Test_Object_Clone(t *testing.T) {
	object := &Object[Action]{}

	object.Set(Action{
		Metadata{
			UID:     "123",
			Version: "123",
		},
		Spec{
			"abc": 123,
			"cde": 123,
		},
	})

	newObj, err := object.Clone()
	if err != nil {
		t.Fatalf("%s", err)
	}
	oldAction := object.Get()
	newAction := newObj.Get()

	if oldAction.UID != newAction.UID ||
		newAction.Version != oldAction.Version ||
		!maps.Equal(newAction.Spec, oldAction.Spec) {
		t.Failed()
	}
}
