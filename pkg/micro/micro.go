package micro

import "net/http"

type MicroService interface {
	// Inject http server handler
	Handle(pattern string, handler http.Handler)
	// need init
	Init() error
	// start server
	Run() error
	// server service uuid
	UUID() string
	// dependency injection
	Inject(m MicroServer)
}

type MicroServer interface {
	Start() error
}
