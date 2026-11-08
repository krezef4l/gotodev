package main

import (
	"fmt"
	"math/rand"
	"sync"
)

func randFunc() <-chan int {
	const N = 100
	a := make(chan int)

	go func() {
		for i := 0; i < N; i++ {
			v := rand.Intn(N) + 1

			a <- v
		}

		close(a)
	}()
	return a
}

func readInt(A <-chan int) <-chan int {
	B := make(chan int)

	go func() {
		for value := range A {
			if value%2 == 0 {
				B <- value
			}
		}
		close(B)
	}()

	return B

}
func printInt(B <-chan int, wg *sync.WaitGroup) {
	defer wg.Done()

	for i := range B {
		fmt.Println("Чётные число:", i)
	}
}

func main() {
	A := randFunc()
	B := readInt(A)

	var wg sync.WaitGroup
	wg.Add(1)
	go printInt(B, &wg)

	wg.Wait()

}

// Есть горутина, которая генерирует случайные числа от 1 до N и отправляет в
// канал A.
// Вторая горутина читает из A и отправляет в канал B только чётные числа.
// Третья горутина читает из B и выводит числа.
// Можно с WaitGroup, можно с <-time.After().
