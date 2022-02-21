package account

import (
	"testing"

	"github.com/ddx2x/flexmeta/pkg/core"
)

func Test_Base_Object_Clone(t *testing.T) {
	object := &core.Object[Account]{}

	object.Set(Account{
		core.Metadata{
			UID:     "123",
			Version: "123",
		},
		Spec{
			AccountType: 3,
		},
	})

	newObj, err := object.Clone()
	if err != nil {
		t.Fatalf("%s", err)
	}
	old := object.Get()
	new := newObj.Get()

	if old.UID != new.UID ||
		new.Version != old.Version {
		t.Failed()
	}

	// newObj.Spec(Spec{
	// 	AccountType: 4,
	// })

	bs, err := newObj.Marshal()
	if err != nil {
		t.Failed()
	}

	_ = bs

	t.Logf("ok")
}
