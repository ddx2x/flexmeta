package core

type Metadata struct {
	Kind    string `json:"kind"`
	Version string `json:"version"`
}

func (m Metadata) GetKind() string           { return m.Kind }
func (m Metadata) SetVersion(version string) { m.Version = version }

type Getter interface {
	GetKind() string
}

type Setter interface {
	SetVersion(version string)
}

type Item interface {
	Getter
	Setter
}

type Object[T Item] struct {
	object T
}

func (o *Object[T]) Clone() *T {
	var new *T
	new = &o.object
	return new
}
