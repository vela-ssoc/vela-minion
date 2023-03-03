package agtapi

import (
	"net/http/pprof"

	"github.com/vela-ssoc/vela-minion/logic/regist"
	"github.com/xgfone/ship/v5"
)

func Debug() regist.Router {
	return new(debugCtrl)
}

type debugCtrl struct{}

func (dc *debugCtrl) Route(arr, _ *ship.RouteGroupBuilder) {
	arr.Route("/debug/pprof").GET(dc.Index)
	arr.Route("/debug/cmdline").GET(dc.Cmdline)
	arr.Route("/debug/profile").GET(dc.Profile)
	arr.Route("/debug/symbol").GET(dc.Symbol).POST(dc.Symbol)
	arr.Route("/debug/trace").GET(dc.Trace)
	arr.Route("/debug/:name").GET(dc.Lookup)
}

func (*debugCtrl) Index(c *ship.Context) error {
	pprof.Index(c.ResponseWriter(), c.Request())
	return nil
}

func (*debugCtrl) Cmdline(c *ship.Context) error {
	pprof.Cmdline(c.ResponseWriter(), c.Request())
	return nil
}

func (*debugCtrl) Profile(c *ship.Context) error {
	pprof.Profile(c.ResponseWriter(), c.Request())
	return nil
}

func (*debugCtrl) Symbol(c *ship.Context) error {
	pprof.Symbol(c.ResponseWriter(), c.Request())
	return nil
}

func (*debugCtrl) Trace(c *ship.Context) error {
	pprof.Trace(c.ResponseWriter(), c.Request())
	return nil
}

func (*debugCtrl) Lookup(c *ship.Context) error {
	name := c.Param("name")
	pprof.Handler(name).ServeHTTP(c.ResponseWriter(), c.Request())
	return nil
}
