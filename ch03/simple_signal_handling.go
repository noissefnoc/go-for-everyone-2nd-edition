package main

import (
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
}
