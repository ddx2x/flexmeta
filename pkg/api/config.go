package api

type Option struct {
	Addr string `json:"addr"`
}

type Options func(*Option)

func WithAddr(addr string) Options {
	return func(o *Option) {
		o.Addr = addr
	}
}
