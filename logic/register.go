package logic

import (
	"github.com/vela-ssoc/vela-minion/logic/agtapi"
	"github.com/vela-ssoc/vela-tunnel"
	"github.com/xgfone/ship/v5"
)

// Register 注册业务逻辑
func Register(tun tunnel.Tunneler, arr, aws *ship.RouteGroupBuilder) {
	// fm FileManager 文件管理模块
	fm := agtapi.FM()
	fm.Route(arr, aws)

	// pprof 模块，用于后期对节点是在线监测
	prof := agtapi.Pprof()
	prof.Route(arr, aws)
}
