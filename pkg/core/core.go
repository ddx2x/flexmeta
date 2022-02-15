package core

type Default struct {
	metadata struct {
		Name    string `json:"name"`
		Kind    string `json:"kind"`
		Version string `json:"version"`
	} `json:"metadata"`
	spec Spec `json:"spec"`
}

// Default.Metadata getter implement
func (d Default) getKind() string    { return d.metadata.Kind }
func (d Default) getVersion() string { return d.metadata.Version }
func (d Default) getSpec() Spec      { return d.spec }

// Default.Metadata setter implement
func (d Default) setVersion(version string) { d.metadata.Version = version }
func (d Default) setSpec(spec Spec)         { d.spec = spec }

type Spec = map[string]interface{}

type getter interface {
	getKind() string
	getVersion() string
	getSpec() Spec
}

type setter interface {
	setVersion(version string)
	setSpec(spec Spec)
}

type Item interface {
	getter
	setter
}

type Object[T Item] struct {
	item T
}

func (o *Object[T]) Kind() string    { return o.item.getKind() }
func (o *Object[T]) Version() string { return o.item.getVersion() }
func (o *Object[T]) Spec() Spec      { return o.item.getSpec() }

func (o *Object[T]) UpdateVersion()       { o.item.setVersion("new version") }
func (o *Object[T]) UpdateSpec(spec Spec) { o.item.setSpec(spec) }

func (o *Object[T]) Clone() *Object[T] {
	var obj Object[T]
	var item = &o.item
	obj.item = *item
	return &obj
}
