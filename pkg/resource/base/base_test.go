package base

import (
	"testing"

	"github.com/ddx2x/flexmeta/pkg/core"
)

func Test_Base_Object_Clone(t *testing.T) {
	object := &core.Object[Base]{}

	object.Set(Base{
		Uid: "123",
	})

	newObj, err := object.Clone()
	if err != nil {
		t.Fatalf("%s", err)
	}
	old := object.Get()
	new := newObj.Get()

	if old.Uid != new.Uid {
		t.Failed()
	}

	// newObj.Spec(Spec{
	// 	UserId: "456",
	// })

	bs, err := newObj.Marshal()
	if err != nil {
		t.Failed()
	}

	_ = bs

	t.Logf("ok")
}
