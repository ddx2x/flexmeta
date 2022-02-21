package api

type Option struct {
	Addr      string `json:"addr"`
	StoreAddr string `json:"store_addr"`
}

type Options func(*Option)

func Addr(addr string) Options {
	return func(o *Option) {
		o.Addr = addr
	}
}

func StoreAddr(addr string) Options {
	return func(o *Option) {
		o.StoreAddr = addr
	}
}
