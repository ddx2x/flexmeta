package core

import (
	"encoding/json"
	"fmt"
)

type Metadata struct {
	Name      string `json:"name"`
	Namespace string `json:"namespace"`
	UID       string `json:"uid"`
	Kind      string `json:"kind"`
	Version   string `json:"version"`
	Area      uint8  `json:"area"`
}

type Spec any

type Objectizable interface {
	~struct {
		Metadata `json:"metadata"`
		Spec     `json:"spec"`
	}
}

type Items[T Objectizable] struct {
	Metadata `json:"metadata"`
	Items    []T `json:"items"`
}

func (i Items[T]) From(objects []Object[T]) Items[T] {
	var md Metadata
	var version string
	var noset bool = false
	for _, obj := range objects {
		item := obj.Get()
		if !noset {
			md = Metadata{
				Kind: fmt.Sprintf("%s%s", item.Kind, "List"),
			}
		}
		if version < item.Version {
			version = item.Version
		}

		i.Items = append(i.Items, item)
	}
	i.Metadata = md
	return i
}

type Object[T Objectizable] struct {
	item T
}

func (o *Object[T]) Set(item T) { o.item = item }

func (o *Object[T]) Spec(spec Spec) { o.item.Spec = spec }

func (o *Object[T]) Marshal() ([]byte, error) { return json.Marshal(o.item) }

func (o *Object[T]) Get() T { return o.item }

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
