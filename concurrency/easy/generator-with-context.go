// Если тема контекста еще не изучена - пропустить задачу
package main

import (
	"fmt"
	"context"
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
				return
			}
		}
	}()
	return out
}

func do3() {
	// cancelCh := make(chan struct{})
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()
	

	generated := generate(ctx, 11)
	for num := range generated {
		fmt.Print(num, " ")
		if num > 14 {
			break
		}
	}
	fmt.Println()
}
