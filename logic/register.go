package logic

import (
	"github.com/vela-ssoc/vela-minion/logic/agtapi"
	"github.com/vela-ssoc/vela-tunnel"
	"github.com/xgfone/ship/v5"
)

func Register(tun tunnel.Tunneler, arr, aws *ship.RouteGroupBuilder) {
	// fm FileManager 文件管理模块
	fm := agtapi.FM()
	fm.Route(arr, aws)
}
