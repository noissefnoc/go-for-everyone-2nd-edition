package main

import (
	"context"
	"fmt"
	"os"
	"os/signal"
	"syscall"
)

func main() {
	defer fmt.Println("done")
	// define signals to handle
	trapSignals := []os.Signal{
		syscall.SIGHUP,
		syscall.SIGINT,
		syscall.SIGTERM,
		syscall.SIGQUIT}
	// create channel for receiving signal
	sigCh := make(chan os.Signal, 1)
	// receive
	signal.Notify(sigCh, trapSignals...)

	// create cancelable context for passing main processing
	ctx, cancel := context.WithCancel(context.Background())
	// wait signal another goroutine
	go func() {
		// block until catch signal
		sig := <-sigCh
		fmt.Println("Got signal", sig)
		cancel()
	}()
	doMain(ctx)
}

func doMain(ctx context.Context) {
	defer fmt.Println("done doMain")
	for {
		select {
		case <-ctx.Done():
			return
		default:
			// do something
		}
		// do something
	}
}
