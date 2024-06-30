package app

import (
	"os"
	"os/signal"
	"sync"
	"syscall"
	"time"
)

func (a *App[acc, prof, regReq, reg]) Start() {
	go func() {
		if err := a.Ginger.Run(); err != nil {
			panic(err)
		}
	}()

	go func() {
		if err := a.GRPC.Run(); err != nil {
			panic(err)
		}
	}()

	done := make(chan struct{})
	go func() {
		c := make(chan os.Signal, 1)
		signal.Notify(c, os.Interrupt, syscall.SIGTERM)
		<-c
		// stop
		wg := new(sync.WaitGroup)
		wg.Add(1)
		go func() {
			a.Logger.WithTrace("exit.ginger").Debugf("stopping...")
			a.Ginger.Shutdown(time.Minute)
			wg.Done()
			a.Logger.WithTrace("exit.ginger").Debugf("stopped")
		}()
		// wait and exit
		wg.Wait()
		close(done)
	}()
	<-done
}
