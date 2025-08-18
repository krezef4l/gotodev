// Если тема контекста еще не изучена - пропустить задачу
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
			// при срабатывании cancel() выполнится этот кейс и мы выйдем из функции
			case <-ctx.Done():
				return
			}
		}
	}()
	return out
}

func main() {
	// добавил контекст с отменой
	ctx, cancel := context.WithCancel(context.Background())

	// передаем теперь контекст, а не канал
	generated := generate(ctx, 11)
	for num := range generated {
		fmt.Print(num, " ")
		if num > 14 {
			// как только наше условее "true" - вызываем cancel()
			cancel()
		}
	}
	fmt.Println()
}
