package agtapi

import (
	"github.com/vela-ssoc/vela-kit/vela"
	"github.com/vela-ssoc/vela-minion/logic/regist"
	"github.com/vela-ssoc/vela-tunnel"
	"github.com/xgfone/ship/v5"
)

func Substance(tun tunnel.Tunneler, env vela.Environment) regist.Router {
	return &substanceCtrl{
		tun: tun,
		env: env,
	}
}

type substanceCtrl struct {
	tun tunnel.Tunneler
	env vela.Environment
}

func (sc *substanceCtrl) Route(arr, _ *ship.RouteGroupBuilder) {
	arr.Route("/substance/sync").GET(sc.Sync)
}

func (sc *substanceCtrl) Sync(c *ship.Context) error {
	return nil
}
