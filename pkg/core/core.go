package core

type Metadata struct {
	Kind    string `json:"kind"`
	Version string `json:"version"`
}

// Metadata Getter implement
func (m Metadata) GetKind() string    { return m.Kind }
func (m Metadata) GetVersion() string { return m.Version }

// Metadata Setter implement
func (m Metadata) SetVersion(version string) { m.Version = version }

type Getter interface {
	GetKind() string
	GetVersion() string
}

type Setter interface {
	SetVersion(version string)
}

type Item interface {
	Getter
	Setter
}

type Object[T Item] struct {
	item T
}

func (o *Object[T]) GetKind() string           { return o.item.GetKind() }
func (o *Object[T]) GetVersion() string        { return o.item.GetVersion() }
func (o *Object[T]) SetVersion(version string) { o.item.SetVersion(version) }

func (o *Object[T]) Clone() *Object[T] {
	var obj Object[T]
	var item = &o.item
	obj.item = *item
	return &obj
}
