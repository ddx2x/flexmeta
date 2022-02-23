package core

import (
	"encoding/json"
)

type IObject interface {
	Unmarshal(a any) error
	Marshal() ([]byte, error)
	
	~struct{} | any
}

type Object[T IObject] struct {
	Item T
}

func NewObject[T IObject](t T) *Object[T] {
	return &Object[T]{}
}

func (o *Object[T]) Set(item T) { o.Item = item }

func (o *Object[T]) Get() T { return o.Item }

func (o *Object[T]) Marshal() ([]byte, error) { return o.Item.Marshal() }

func (o *Object[T]) Unmarshal(a any) error { return o.Item.Unmarshal(a) }

func (o *Object[T]) From(i any) error {
	b, err := json.Marshal(i)
	if err != nil {
		return err
	}
	return o.Unmarshal(b)
}

func (o *Object[T]) ToMap() (map[string]any, error) {
	bs, err := json.Marshal(&o.Item)
	if err != nil {
		return nil, err
	}
	var r map[string]any
	if err = json.Unmarshal(bs, r); err != nil {
		return nil, err
	}
	return r, nil
}

func (o *Object[T]) Clone() (*Object[T], error) {
	var obj Object[T]
	var target T
	src, err := json.Marshal(&o.Item)
	if err != nil {
		return nil, err
	}
	if err := json.Unmarshal(src, &target); err != nil {
		return nil, err
	}
	obj.Item = target
	return &obj, nil
}

type EventType = string

const (
	ADDED    EventType = "ADDED"
	MODIFIED EventType = "MODIFIED"
	DELETED  EventType = "DELETED"
	REMOVED  EventType = "REMOVED"
)

type Event struct {
	Type   EventType `json:"type"`
	Object any       `json:"object"`
}
