package main

import (
	"context"
	"fmt"
)

// Есть функция generate(), которая генерит числа
// Функция использует канал отмены. Переделать на контекст.
func generate(ctx context.Context, start int) <-chan int {
	out := make(chan int)
	go func() {
		defer close(out)
		for i := start; ; i++ {
			select {
			case out <- i:
			case <-ctx.Done():
				fmt.Print("Отмена")
				return
			}
		}
	}()
	return out
}

func main() {
	ctx, cancel := context.WithCancel(context.Background())

	generated := generate(ctx, 11)
	for num := range generated {
		fmt.Print(num, " ")
		if num > 14 {
			cancel()
		}
	}
	fmt.Println()
}
