package regist

import "github.com/xgfone/ship/v5"

type Router interface {
	Route(arr, aws *ship.RouteGroupBuilder)
}
