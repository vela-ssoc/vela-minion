package launch

import (
	"context"
	"time"

	"github.com/vela-ssoc/backend-common/logback"
	"github.com/vela-ssoc/backend-common/netutil"
	"github.com/vela-ssoc/backend-common/validate"
	"github.com/vela-ssoc/vela-minion/logic"
	"github.com/vela-ssoc/vela-tunnel"
	"github.com/xgfone/ship/v5"
)

func Run(parent context.Context, hide tunnel.Hide, slog logback.Logger) error {
	ctx, cancel := context.WithCancel(parent)
	defer cancel()

	slog.Info("准备连接服务端")
	opts := []tunnel.Option{tunnel.WithLogger(slog), tunnel.WithInterval(10 * time.Minute)}
	tun, err := tunnel.Dial(ctx, hide, opts...)
	if err != nil {
		return err
	}
	slog.Infof("连接服务端 %s 成功", tun.BrkAddr())

	name := tun.NodeName()
	han := ship.Default()
	han.Logger = slog
	han.NotFound = netutil.Notfound(name)
	han.HandleError = netutil.ErrorFunc(name)
	han.Validator = validate.New()

	group := han.Group("/api")
	arr := group.Group("/arr")
	aws := group.Group("/aws")
	logic.Register(tun, arr, aws) // 注册业务逻辑

	che := make(chan error, 1)
	dt := &daemonTun{
		tun:  tun,
		han:  han,
		slog: slog,
		ctx:  ctx,
		err:  che,
	}
	go dt.Run()

	select {
	case <-ctx.Done():
	case err = <-che:
	}

	return err
}
