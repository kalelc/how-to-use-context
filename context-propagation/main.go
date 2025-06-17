package main

import (
	"context"
	"fmt"
	"time"
)

func main() {
	runContext()
}

func runContext() {
	ctx, cancel := context.WithTimeout(context.Background(), 2*time.Second)
	defer cancel()

	err := ControllerWithGoodContext(ctx)
	if err != nil {
		fmt.Println("Result:", err)
	} else {
		fmt.Println("Result: success")
	}
}

func ControllerWithGoodContext(ctx context.Context) error {
	fmt.Println("Controller: Calling service...")

	return ServiceWithGoodContext(ctx)
}

func ServiceWithGoodContext(ctx context.Context) error {
	fmt.Println("Service: Doing work...")

	select {
	case <-time.After(3 * time.Second): // Simulates long task
		fmt.Println("Service: Finished")
		return nil
	case <-ctx.Done():
		fmt.Println("Service: Canceled due to:", ctx.Err())
		return ctx.Err()
	}
}
