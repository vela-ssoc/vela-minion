package launch

import (
	"context"
	"net/http"

	"github.com/vela-ssoc/backend-common/logback"
	"github.com/vela-ssoc/vela-tunnel"
)

type daemonTun struct {
	tun  tunnel.Tunneler
	han  http.Handler
	slog logback.Logger
	ctx  context.Context
	err  chan<- error
}

func (dt *daemonTun) Run() {
	for {
		srv := &http.Server{Handler: dt.han}
		err := srv.Serve(dt.tun.Listen())
		dt.slog.Warnf("与服务端断开连接：%v", err)
		if err = dt.ctx.Err(); err != nil {
			dt.err <- err
			return
		}
		if err = dt.tun.Reconnect(nil); err != nil {
			dt.err <- err
			return
		}
		dt.slog.Info("与服务端重连成功")
	}
}
