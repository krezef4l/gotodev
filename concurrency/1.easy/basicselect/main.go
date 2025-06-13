package main

import (
	"context"
	"fmt"
	"time"
)

// Что выведется?
// Развернуто объяснить почему.

func main() {
	timeout := 3 * time.Second
	ctx, cancel := context.WithTimeout(context.Background(), timeout)

	for {
		select {
		case <-time.After(1 * time.Second):
			time.Sleep(5 * time.Millisecond)
			fmt.Println("waited for 1 sec")
		case <-time.After(2 * time.Second):
			fmt.Println("waited for 2 sec")
			cancel()
		case <-time.After(3 * time.Second):
			fmt.Println("waited for 3 sec")
		case <-ctx.Done():
			fmt.Println(ctx.Err())
			return
		}
	}
}
