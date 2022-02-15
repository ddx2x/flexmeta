package core

import (
	"fmt"
	"testing"
)

type TestObject struct {
	Default
}

func Test_Object_Clone(t *testing.T) {
	object := &Object[TestObject]{}
	object.UpdateVersion()
	newObj := object.Clone()

	if fmt.Sprintf("%v#", object) == fmt.Sprintf("%v#", newObj) ||
		object.Version() != newObj.Version() {
		t.Failed()
	}

	object.UpdateSpec(map[string]interface{}{"a": 123})
	new2Obj := object.Clone()
	value := new2Obj.Spec()

	_ = value
}
