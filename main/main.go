package main

import (
	"context"
	"os"
	"os/signal"
	"syscall"

	"github.com/vela-ssoc/backend-common/encipher"
	"github.com/vela-ssoc/backend-common/logback"
	"github.com/vela-ssoc/vela-minion/banner"
	"github.com/vela-ssoc/vela-minion/launch"
	"github.com/vela-ssoc/vela-tunnel"
)

func main() {
	banner.Print(os.Stdout)
	slog := logback.Stdout()

	var hide tunnel.Hide
	if err := encipher.ReadFile(os.Args[0], &hide); err != nil {
		slog.Errorf("读取 hide 数据错误：%v", err)
		return
	}

	// 监听停止信号
	cares := []os.Signal{syscall.SIGTERM, syscall.SIGHUP, syscall.SIGKILL, syscall.SIGINT}
	ctx, cancel := signal.NotifyContext(context.Background(), cares...)
	defer cancel()
	slog.Info("按 Ctrl+C 结束运行")

	err := launch.Run(ctx, hide, slog)
	slog.Warn("程序已结束运行：%v", err)
}
