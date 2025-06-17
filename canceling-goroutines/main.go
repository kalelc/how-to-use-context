package main

import (
	"context"
	"fmt"
	"time"
)

func main() {
	runWithContext()
}

func worker(ctx context.Context) {
	for {
		select {
		case <-ctx.Done():
			fmt.Println("Worker stopped:", ctx.Err())
			return
		default:
			fmt.Println("Working...")
			time.Sleep(500 * time.Millisecond)
		}
	}
}

func runWithContext() {
	ctx, cancel := context.WithCancel(context.Background())
	go worker(ctx)

	time.Sleep(2 * time.Second)
	cancel() // stops the worker
	time.Sleep(1 * time.Second)
}
