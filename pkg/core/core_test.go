package core

import (
	"testing"
)

var _ IObject = &Action{}

type Action struct {
	Uid     string `json:"uid" bson:"_id"`
	Version string `json:"version" bson:"version"`
	Id      string
}

func (a Action) Unmarshal(i any) error {
	return nil
}

func (a Action) Marshal() ([]byte, error) {
	return nil, nil
}

func Test_Object_Clone(t *testing.T) {
	object := &Object[Action]{}

	object.Set(Action{
		Uid:     "123",
		Version: "123",
		Id:      "123",
	})

	newObj, err := object.Clone()
	if err != nil {
		t.Fatalf("%s", err)
	}
	old := object.Get()
	new := newObj.Get()

	if old.Uid != new.Uid || old.Version != new.Version {
		t.Failed()
	}

	t.Logf("ok")
}
