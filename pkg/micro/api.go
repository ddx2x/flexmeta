package micro

import (
	"net/http"
)

type Api struct {
	svc MicroService
}

func (m *Api) Inject(handler http.Handler) error {
	if err := m.svc.Init(); err != nil {
		return err
	}
	m.svc.Handle("/", handler)
	return nil
}

func (m *Api) Start() error { return m.svc.Run() }

func NewService(svc MicroService) *Api {
	return &Api{
		svc: svc,
	}
}
