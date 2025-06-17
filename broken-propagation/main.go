package main

import (
	"context"
	"fmt"
	"time"
)

func main() {
	runWithContext()
}

func runWithContext() {
	ctx, cancel := context.WithTimeout(context.Background(), 2*time.Second)
	defer cancel()

	err := ControllerWithBadContext(ctx)
	if err != nil {
		fmt.Println("Result:", err)
	} else {
		fmt.Println("Result: success")
	}
}

func ControllerWithBadContext(ctx context.Context) error {
	fmt.Println("Controller: Calling service...")

	return ServiceWithBadContext(ctx)
}

func ServiceWithBadContext(_ context.Context) error {
	fmt.Println("Service: Doing work...")

	ctx := context.Background()

	select {
	case <-time.After(3 * time.Second):
		fmt.Println("Service: Finished")
		return nil
	case <-ctx.Done():
		fmt.Println("Service: Canceled due to:", ctx.Err())
		return ctx.Err()
	}
}
