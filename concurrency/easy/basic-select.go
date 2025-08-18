package main

import (
	"context"
	"fmt"
	"time"
)

// Что выведет?

// Ответ:

// waited for 1 sec
// waited for 1 sec
// третий раз наверное не успеет выполниться, так как контекст закончится раньше, потому что опереации, которые не входят в select тоже займут какое-то небольшое время и третья "wait for 1 sec" не успеет и будет:
// deadline exceeded

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
