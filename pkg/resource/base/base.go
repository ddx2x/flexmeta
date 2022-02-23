package base
type Base struct {
	Uid string `json:"uid" bson:"_id"`
}

func (b Base) Unmarshal(i any) error {
	return nil
}

func (b Base) Marshal() ([]byte, error) {
	return nil, nil
}
