package common

import (
	"fmt"
	"os"
	"os/signal"
	"syscall"
)

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
			return
		default:
			fmt.Println("other signal", s)
		}
	}
}
