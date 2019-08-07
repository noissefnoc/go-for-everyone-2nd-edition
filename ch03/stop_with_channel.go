package main

import (
	"fmt"
	"sync"
)

var wg sync.WaitGroup

func main() {
	queue := make(chan string)
	for i := 0; i < 2; i++ { // generate 2 goroutine
		wg.Add(1)
		go fetchURLWithCh(queue)
	}

	queue<- "https://www.example.com"
	queue<- "https://www.example.net"
	queue<- "https://www.example.net/foo"
	queue<- "https://www.example.net/bar"

	close(queue) // end message to goroutine
	wg.Wait()    // wait for all goroutines
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
			wg.Done()
			return
		}
	}
}
