package core

import (
	"encoding/json"
)

type IObject interface {
	~struct {
		UID     string `json:"uid"`
		Kind    string `json:"kind"`
		Version string `json:"version"`
	}
	UpdateVersion()
}

type BObject struct {
	UID     string `json:"uid"`
	Kind    string `json:"kind"`
	Version string `json:"version"`
}

type X map[any]any

func (b BObject) UpdateVersion() { b.Version = "new" }

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

func (o *Object[T]) Get() T { return o.Item }

func (o *Object[T]) Marshal() ([]byte, error) { return json.Marshal(o.Item) }

func (o *Object[T]) Unmarshal(b []byte) error { return json.Unmarshal(b, &o.Item) }

func (o *Object[T]) From(i interface{}) error {
	b, err := json.Marshal(i)
	if err != nil {
		return err
	}
	return o.Unmarshal(b)
}

func (o *Object[T]) ToMap() (map[string]interface{}, error) {
	bs, err := json.Marshal(&o.Item)
	if err != nil {
		return nil, err
	}
	var r map[string]interface{}
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
