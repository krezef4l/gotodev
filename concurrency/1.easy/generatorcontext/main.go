package main

import (
	"context"
	"fmt"
)

// Есть функция generate(), которая генерит числа. Функция использует канал
// отмены. Переделать на контекст: 'Done!' в 18-ой строке должно быть выведено
// на экран.

func generate(ctx context.Context, start int) (<-chan int, <-chan struct{}) {
	out := make(chan int)
	done := make(chan struct{})
	go func() {
		defer close(done)
		defer close(out)
		for i := start; ; i++ {
			select {
			case out <- i:
			case <-ctx.Done():
				fmt.Println("Done!")
				return
			}
		}
	}()
	return out, done
}

func main() {
	ctx, cancel := context.WithCancel(context.Background())

	generated, done := generate(ctx, 11)
	for num := range generated {
		fmt.Print(num, " ")
		if num > 14 {
			cancel()
			break
		}
	}
	<-done
	fmt.Println()
}
