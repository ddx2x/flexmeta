package micro

import (
	"net/http"
)

type Api[S MicroService] struct {
	svc S
}

func (m *Api[S]) Inject(handler http.Handler) error {
	if err := m.svc.Init(); err != nil {
		return err
	}
	m.svc.Handle("/", handler)
	return nil
}

func (m *Api[S]) Start() error { return m.svc.Run() }

func NewService[S MicroService](svc S) *Api[S] {
	return &Api[S]{
		svc: svc,
	}
}
