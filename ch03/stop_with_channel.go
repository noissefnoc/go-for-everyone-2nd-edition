package main

import (
	"fmt"
	"sync"
)

var wgc sync.WaitGroup

func main() {
	queue := make(chan string)
	for i := 0; i < 2; i++ { // generate 2 goroutine
		wgc.Add(1)
		go fetchURLWithCh(queue)
	}

	queue <- "https://www.example.com"
	queue <- "https://www.example.net"
	queue <- "https://www.example.net/foo"
	queue <- "https://www.example.net/bar"

	close(queue) // end message to goroutine
	wgc.Wait()   // wait for all goroutines
}

func fetchURLWithCh(queue chan string) {
	for {
		url, more := <-queue // if close then more is false
		if more {
			// url get result
			fmt.Println("fetching", url)
			// some processing
		} else {
			fmt.Println("worker exit")
			wgc.Done()
			return
		}
	}
}
