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
	newObj := object.Clone()
	fmt.Println(newObj)
}
