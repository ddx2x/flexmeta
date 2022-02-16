package core

import "encoding/json"

var _ Getter = &Metadata{}

type Metadata struct {
	UID     string `json:"uid"`
	Version string `json:"version"`
}

func (m Metadata) Get() Metadata { return m }

type Spec map[string]any

type Getter interface {
	Get() Metadata
}

type IObject interface {
	~struct {
		Metadata `json:"metadata"`
		Spec     `json:"spec"`
	}
	Getter
}

type Object[T IObject] struct {
	item T
}

func (o *Object[T]) Set(item T)         { o.item = item }
func (o *Object[T]) Metadata() Metadata { return o.item.Get() }

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
