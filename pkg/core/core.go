package core

import "encoding/json"

type Metadata struct {
	UID     string `json:"uid"`
	Version string `json:"version"`
}

type Spec map[string]any

type IObject interface {
	~struct {
		Metadata `json:"metadata"`
		Spec     `json:"spec"`
	}
}

type Object[T IObject] struct {
	item T
}

func (o *Object[T]) Set(item T) { o.item = item }
func (o *Object[T]) Get() T     { return o.item }

func (o *Object[T]) Clone() (*Object[T], error) {
	var obj Object[T]
	var target T
	src, err := json.Marshal(&o.item)
	if err != nil {
		return nil, err
	}
	if err := json.Unmarshal(src, &target); err != nil {
		return nil, err
	}
	obj.item = target
	return &obj, nil
}
