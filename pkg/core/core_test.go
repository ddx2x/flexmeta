package core

import (
	"fmt"
	"testing"
)

type TestObject struct {
	Metadata
}

func Test_Object_Clone(t *testing.T) {
	object := &Object[TestObject]{}
	object.item.SetVersion("123")
	newObj := object.Clone()

	if fmt.Sprintf("%x", object) == fmt.Sprintf("%x", newObj) ||
		object.GetVersion() != newObj.GetVersion() {
		t.Failed()
	}
}
