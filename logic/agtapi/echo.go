package agtapi

import (
	"github.com/gorilla/websocket"
	"github.com/vela-ssoc/backend-common/netutil"
	"github.com/vela-ssoc/vela-minion/logic/regist"
	"github.com/vela-ssoc/vela-tunnel"
	"github.com/xgfone/ship/v5"
)

func Echo(tun tunnel.Tunneler) regist.Router {
	name := tun.NodeName()
	upg := netutil.Upgrade(name)
	return &echoCtrl{
		name: name,
		upg:  upg,
	}
}

type echoCtrl struct {
	name string
	upg  websocket.Upgrader
}

func (ec *echoCtrl) Route(_, aws *ship.RouteGroupBuilder) {
	aws.Route("/echo").GET(ec.Echo)
}

func (ec *echoCtrl) Echo(c *ship.Context) error {
	w, r := c.ResponseWriter(), c.Request()
	ws, err := ec.upg.Upgrade(w, r, nil)
	if err != nil {
		return err
	}
	//goland:noinspection GoUnhandledErrorResult
	defer ws.Close()

	for {
		mt, dat, err := ws.ReadMessage()
		if err != nil {
			c.Warnf("websocket 读取错误：%s", err)
			break
		}

		c.Infof("%s 收到了消息：%s", ec.name, dat)

		if err = ws.WriteMessage(mt, dat); err != nil {
			c.Warnf("websocket 回写消息错误：%s", err)
			break
		}
	}

	c.Warnf("socket 连接已经断开：%v", err)

	return nil
}
