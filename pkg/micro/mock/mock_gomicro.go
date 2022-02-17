package mock

import (
	"net/http"
	"testing"

	mi "github.com/ddx2x/flexmeta/pkg/micro"
	"github.com/micro/micro/v2/cmd"
)

type mockgw struct{}

func (g mockgw) Handle(pattern string, handler http.Handler) {}

func (g mockgw) Init() error { return nil }

func (g mockgw) UUID() string { return "" }

func (g mockgw) Run() error { cmd.Init(); return nil }

func (g mockgw) Inject(m mi.MicroServer) {}

func Test_go_microgw_mock(t *testing.T) {

}

type mockapi struct{}

func (g mockapi) Handle(pattern string, handler http.Handler) {}

func (g mockapi) Init() error { return nil }

func (g mockapi) UUID() string { return "" }

func (g mockapi) Run() error { cmd.Init(); return nil }

func (g mockapi) Inject(m mi.MicroServer) {}

func Test_go_microapi_mock(t *testing.T) {

}
