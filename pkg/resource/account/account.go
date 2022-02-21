package account

import (
	"github.com/ddx2x/flexmeta/pkg/core"
)

type Spec struct {
	Password    string `json:"password" bson:"password"`
	AccountType uint   `json:"account_type" bson:"account_type"`
}

type Account struct {
	core.Metadata `json:"metadata"`
	Spec          `json:"spec"`
}
