package core

import (
	"testing"
)

type Action struct {
	Metadata `json:"metadata"`
	Spec     `json:"spec"`
}

func Test_Object_Clone(t *testing.T) {
	object := &Object[Action]{}
	action := Action{
		Metadata{
			UID: "123",
		},
		Spec{},
	}
	object.Set(action)
	newObj, err := object.Clone()
	if err != nil {
		t.Fatalf("%s", err)
	}
	newAction := newObj.Metadata()
	if newAction.UID != action.UID {
		t.Failed()
	}
}
