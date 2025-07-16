package base

import (
	"log/slog"
	"os"
	"os/signal"
	"syscall"
)

func GetSignalChan(sign ...os.Signal) chan os.Signal {
	c := make(chan os.Signal)
	//监听指定信号 ctrl+c kill
	if len(sign) > 0 {
		signal.Notify(c, sign...)
	} else {
		signal.Notify(c, syscall.SIGHUP, syscall.SIGINT, syscall.SIGTERM,
			syscall.SIGQUIT)
	}
	return c
}

func StartSignalLister(f func()) {
	c := make(chan os.Signal)
	//监听指定信号 ctrl+c kill
	signal.Notify(c, syscall.SIGHUP, syscall.SIGINT, syscall.SIGTERM,
		syscall.SIGQUIT)
	SignalHandler(c, f)
	return
}

func SignalHandler(c chan os.Signal, f func()) {
	for s := range c {
		switch s {
		case syscall.SIGHUP, syscall.SIGINT, syscall.SIGTERM, syscall.SIGQUIT:
			f()
			slog.Info("SignalHandler Over", "sign", s.String())
			return
		default:
			slog.Info("other signal", "sign", s.String())
		}
	}
}
