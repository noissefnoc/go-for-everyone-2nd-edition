package main

import (
	"log"
	"os"
	"os/signal"
	"syscall"
	"time"
)

type MySignal struct {
	message string
}

func (s MySignal) String() string {
	return s.message
}

func (s MySignal) Signal() {}

func main() {
	log.Println("[info] Start")
	trapSignals := []os.Signal{
		syscall.SIGHUP,
		syscall.SIGINT,
		syscall.SIGTERM,
		syscall.SIGQUIT}

	sigCh := make(chan os.Signal, 1)

	time.AfterFunc(10*time.Second, func() {
		sigCh<-MySignal{"timed out"}
	})

	signal.Notify(sigCh, trapSignals...)

	sig := <-sigCh
	switch s := sig.(type) {
	case syscall.Signal:
		log.Printf("[info] Got signal: %s(%d)", s, s)
	case MySignal:
		log.Printf("[info] %s", s)
	}
}
