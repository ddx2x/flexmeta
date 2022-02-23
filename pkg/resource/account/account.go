package account

type Account struct {
	Uid string `json:"uid" bson:"_id"`
}

func (a Account) Unmarshal(i any) error {
	return nil
}

func (a Account) Marshal() ([]byte, error) {
	return nil, nil
}
