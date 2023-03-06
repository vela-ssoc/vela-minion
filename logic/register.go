package logic

import (
	"github.com/vela-ssoc/vela-kit/vela"
	"github.com/vela-ssoc/vela-minion/logic/agtapi"
	"github.com/vela-ssoc/vela-tunnel"
	"github.com/xgfone/ship/v5"
)

// Register 注册业务逻辑
func Register(tun tunnel.Tunneler, env vela.Environment, arr, aws *ship.RouteGroupBuilder) {
	agtapi.FM().Route(arr, aws)      // fm FileManager 文件管理模块
	agtapi.Pprof().Route(arr, aws)   // pprof 模块，用于运行时对节点是在线监测
	agtapi.Echo(tun).Route(arr, aws) // 简单的双向流响应 echo
	agtapi.Substance(tun, env).Route(arr, aws)
}
