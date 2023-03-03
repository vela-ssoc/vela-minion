package main

import (
	"context"
	"os"
	"os/signal"
	"syscall"

	"github.com/vela-ssoc/backend-common/logback"
	"github.com/vela-ssoc/vela-minion/banner"
	"github.com/vela-ssoc/vela-minion/launch"
	"github.com/vela-ssoc/vela-tunnel"
)

func main() {
	banner.Print(os.Stdout)
	slog := logback.Stdout()

	addr := tunnel.Addresses{
		{TLS: true, Addr: "172.31.61.168", Name: "local.eastmoney.com"},
		{Addr: "172.31.61.168:8180"},
	}
	hide := tunnel.Hide{
		Semver:   "0.0.1-delve",
		Ethernet: addr,
		Internet: nil,
	}

	// 监听停止信号
	cares := []os.Signal{syscall.SIGTERM, syscall.SIGHUP, syscall.SIGKILL, syscall.SIGINT}
	ctx, cancel := signal.NotifyContext(context.Background(), cares...)
	defer cancel()
	slog.Info("按 Ctrl+C 结束运行")

	err := launch.Run(ctx, hide, slog)
	slog.Warn("程序已结束运行：%v", err)
}
