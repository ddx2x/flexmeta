package micro

import "net/http"

type InterceptType uint8

const (
	SelfHandle InterceptType = iota
	Redirect
	Interrupt
	NotAuthorized
	Next
)

type Handler func(http.Handler) http.Handler

type Intercept func(w http.ResponseWriter, r *http.Request) InterceptType

type Gateway[S MicroService] struct {
	svc S
}

func NewGateway[S MicroService](svc S) *Gateway[S] {
	return &Gateway[S]{
		svc: svc,
	}
}

func (g *Gateway[S]) Start() error { 
	if err := g.svc.Init(); err != nil {
		return err
	}
	return g.svc.Run() 
}

func (g *Gateway[S]) Intercept(self http.Handler, its ...Intercept) Handler {
	return func(redirect http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			state := Redirect
			for _, intercept := range its {
				state = intercept(w, r)
				if state != Next {
					break
				}
			}
			switch state {
			case SelfHandle:
				self.ServeHTTP(w, r)
			case Interrupt:
				w.WriteHeader(http.StatusNotFound)
				return
			case NotAuthorized:
				w.WriteHeader(http.StatusUnauthorized)
				_, _ = w.Write([]byte("access denied"))
				return
			default:
				redirect.ServeHTTP(w, r)
			}
		})
	}
}
