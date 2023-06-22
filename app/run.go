package app

import (
	"os"
	"os/signal"
	"syscall"
)

func (a *app) Start() {
	go func() {
		if err := a.ginger.Run(); err != nil {
			panic(err)
		}
	}()

	done := make(chan struct{})
	go func() {
		c := make(chan os.Signal, 1)
		signal.Notify(c, os.Interrupt, syscall.SIGTERM)
		<-c
		close(done)
	}()
	<-done
}
