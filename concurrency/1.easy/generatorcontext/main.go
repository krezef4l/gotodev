package main

import (
	"fmt"
)

// Есть функция generate(), которая генерит числа. Функция использует канал отмены. Переделать на контекст.

func generate(cancel <-chan struct{}, start int) <-chan int {
	out := make(chan int)
	go func() {
		defer close(out)
		for i := start; ; i++ {
			select {
			case out <- i:
			case <-cancel:
				return
			}
		}
	}()
	return out
}

func main() {
	cancelCh := make(chan struct{})

	generated := generate(cancelCh, 11)
	for num := range generated {
		fmt.Print(num, " ")
		if num > 14 {
			break
		}
	}
	fmt.Println()
}
