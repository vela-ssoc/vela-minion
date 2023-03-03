package agtapi

import (
	"net/http/pprof"

	"github.com/vela-ssoc/vela-minion/logic/regist"
	"github.com/xgfone/ship/v5"
)

func Pprof() regist.Router {
	return new(pprofCtrl)
}

type pprofCtrl struct{}

func (pc *pprofCtrl) Route(arr, _ *ship.RouteGroupBuilder) {
	arr.Route("/debug/pprof").GET(pc.Index)
	arr.Route("/debug/cmdline").GET(pc.Cmdline)
	arr.Route("/debug/profile").GET(pc.Profile)
	arr.Route("/debug/symbol").GET(pc.Symbol).POST(pc.Symbol)
	arr.Route("/debug/trace").GET(pc.Trace)
	arr.Route("/debug/:name").GET(pc.Lookup)
}

func (*pprofCtrl) Index(c *ship.Context) error {
	pprof.Index(c.ResponseWriter(), c.Request())
	return nil
}

func (*pprofCtrl) Cmdline(c *ship.Context) error {
	pprof.Cmdline(c.ResponseWriter(), c.Request())
	return nil
}

func (*pprofCtrl) Profile(c *ship.Context) error {
	pprof.Profile(c.ResponseWriter(), c.Request())
	return nil
}

func (*pprofCtrl) Symbol(c *ship.Context) error {
	pprof.Symbol(c.ResponseWriter(), c.Request())
	return nil
}

func (*pprofCtrl) Trace(c *ship.Context) error {
	pprof.Trace(c.ResponseWriter(), c.Request())
	return nil
}

func (*pprofCtrl) Lookup(c *ship.Context) error {
	name := c.Param("name")
	pprof.Handler(name).ServeHTTP(c.ResponseWriter(), c.Request())
	return nil
}
