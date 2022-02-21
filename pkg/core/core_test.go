package core

import (
	"testing"
)

type Action struct {
	Metadata `json:"metadata"`
	Spec     any `json:"spec"`
}

type ActionSpec struct {
	Id string
}

func Test_Object_Clone(t *testing.T) {
	object := &Object[Action]{}

	object.Set(Action{
		Metadata{
			UID:     "123",
			Version: "123",
		},
		ActionSpec{
			Id: "123",
		},
	})

	newObj, err := object.Clone()
	if err != nil {
		t.Fatalf("%s", err)
	}
	old := object.Get()
	new := newObj.Get()

	if old.UID != new.UID || old.Version != new.Version {
		t.Failed()
	}

	t.Logf("ok")
}
