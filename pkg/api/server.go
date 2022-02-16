package api

import "net/http"

type Server[T http.Server] struct {
	serv T
}
