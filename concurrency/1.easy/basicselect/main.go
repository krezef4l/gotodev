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

// Здесь работает только первая ветка селекта!
// Дано 3 секунды до срабатывания контекста, итерация +-1.005сек
// 2 итерации и в конце сработает контекст, а не cancel
