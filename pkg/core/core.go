package core

import (
	"encoding/json"
)

type Metadata struct {
	Name      string `json:"name"`
	Namespace string `json:"namespace"`
	UID       string `json:"uid"`
	Kind      string `json:"kind"`
	Version   string `json:"version"`
	Area      uint8  `json:"area"`
}
type Objectizable interface {
	~struct {
		Metadata `json:"metadata"`
	} | any
}

type Items[T any] struct {
	Metadata `json:"metadata"`
	Items    []T `json:"Items"`
}

func (i Items[T]) From(objects []Object[T]) Items[T] {
	// var md Metadata
	// var version string
	// var noset bool = false
	// for _, obj := range objects {
	// 	Item := obj.Get()
	// 	if !noset {
	// 		md = Metadata{
	// 			Kind: fmt.Sprintf("%s%s", Item.Kind, "List"),
	// 		}
	// 	}
	// 	if version < Item.Version {
	// 		version = Item.Version
	// 	}

	// 	i.Items = append(i.Items, Item)
	// }
	// i.Metadata = md
	return i
}

type Object[T any] struct {
	Item T
}

func (o *Object[T]) Set(item T) { o.Item = item }

func (o *Object[T]) Marshal() ([]byte, error) { return json.Marshal(o.Item) }

func (o *Object[T]) Get() T { return o.Item }

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
