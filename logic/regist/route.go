package regist

import "github.com/xgfone/ship/v5"

// Router 路由模块
type Router interface {
	// Route 注册路由
	// arr: Agent Request Response 请求响应模式的接口
	// aws: Agent WebSocket 流式传输接口
	Route(arr, aws *ship.RouteGroupBuilder)
}
