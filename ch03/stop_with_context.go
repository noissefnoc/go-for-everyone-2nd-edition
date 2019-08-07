package main

import (
	"context"
	"fmt"
	"sync"
)

var wg sync.WaitGroup

func main() {
	// generate context for cancel
	ctx, cancel := context.WithCancel(context.Background())
	queue := make(chan string)
	for i := 0; i < 2; i++ {
		wg.Add(1)
		go fetchURLWithCtx(ctx, queue)
	}

	queue<- "https://www.example.com"
	queue<- "https://www.example.net"
	queue<- "https://www.example.net/foo"
	queue<- "https://www.example.net/bar"

	cancel()  // end ctx
	wg.Wait() // wait for all goroutine ends
}

func fetchURLWithCtx(ctx context.Context, queue chan string) {
	for {
		select {
		case <-ctx.Done():
			fmt.Println("worker exit")
			wg.Done()
			return
		case url := <-queue:
			// fetch URL
			fmt.Println("fetching", url)
			// do some professing
		}
	}
}
