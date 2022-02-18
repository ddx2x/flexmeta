package resource

import (
	"testing"

	"github.com/ddx2x/flexmeta/pkg/core"
)

func Test_Base_Object_Clone(t *testing.T) {
	object := &core.Object[Base]{}

	object.Set(&Base{
		core.Metadata{
			UID:     "123",
			Version: "123",
		},
		BaseSpec{
			UserId: "123",
		},
	})

	newObj, err := object.Clone()
	if err != nil {
		t.Fatalf("%s", err)
	}
	oldAction := object.Get()
	newAction := newObj.Get()

	if oldAction.UID != newAction.UID ||
		newAction.Version != oldAction.Version {
		t.Failed()
	}

	t.Logf("ok")
}
