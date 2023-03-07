package agtapi

import (
	"context"
	"fmt"
	"io"
	"net/http"
	"sync"
	"time"

	"github.com/gorilla/websocket"
	"github.com/vela-ssoc/backend-common/netutil"
	"github.com/vela-ssoc/vela-minion/logic/regist"
	"github.com/vela-ssoc/vela-tunnel"
	"github.com/xgfone/ship/v5"
)

// Echo 简单的流式输入输出 demo
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

func (ec *echoCtrl) Route(arr, aws *ship.RouteGroupBuilder) {
	aws.Route("/echo").GET(ec.Echo)
	arr.Route("/sse").GET(ec.SSE)
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

func (ec *echoCtrl) SSE(c *ship.Context) error {
	sse, err := ec.newSSE(c)
	if err != nil {
		return err
	}
	time.Sleep(10 * time.Second)
	sse.Sent("你好呀")
	time.Sleep(10 * time.Second)

	return nil
}

func (*echoCtrl) newSSE(c *ship.Context) (*sentEvents, error) {
	acc := c.GetReqHeader(ship.HeaderAccept)
	if acc != "text/event-stream" {
		return nil, ship.ErrUnsupportedMediaType
	}
	w, r := c.ResponseWriter(), c.Request()
	fls, ok := w.(http.Flusher)
	if !ok {
		return nil, ship.ErrUnsupportedMediaType
	}
	w.Header().Set(ship.HeaderContentType, "text/event-stream")
	w.Header().Set(ship.HeaderCacheControl, "no-cache")
	w.Header().Set(ship.HeaderConnection, "keep-alive")
	w.Header().Set(ship.HeaderAccessControlAllowOrigin, "*")
	w.WriteHeader(http.StatusOK)
	fls.Flush()

	se := &sentEvents{
		fls: fls,
		wtr: w,
		ctx: r.Context(),
	}

	return se, nil
}

type sentEvents struct {
	fls   http.Flusher
	wtr   io.Writer
	uid   uint64
	mutex sync.Mutex
	ctx   context.Context
}

func (se *sentEvents) Sent(data string) error {
	se.mutex.Lock()
	defer se.mutex.Unlock()

	if err := se.ctx.Err(); err != nil {
		return err
	}
	se.uid++

	if _, err := fmt.Fprintf(se.wtr, "id: %d\ndata:%s\n\n", se.uid, data); err != nil {
		return err
	}
	se.fls.Flush()
	return nil
}
