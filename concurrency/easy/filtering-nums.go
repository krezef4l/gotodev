// Есть горутина, которая генерирует случайные числа от 1 до N и отправляет в канал A.
// Вторая горутина читает из A и отправляет в канал B только чётные числа.
// Третья горутина читает из B и выводит числа.
// Можно с WaitGroup, можно с <-time.After()

package main

import (
	"fmt"
	"math/rand"
	"sync"
)

func gen(c chan<- int, N int) {
	const min = 1
	rnd := rand.Intn(N-min) + min
	c <- rnd
}

func filter(in <-chan int, out chan<- int) {
	num := <-in

	if num%2 == 0 {
		out <- num
	}

	close(out)
}

func print(c <-chan int) {
	fmt.Println(<-c)
}

func main() {
	var wg sync.WaitGroup

	A := make(chan int)
	B := make(chan int)

	wg.Add(3)
	go func() {
		defer wg.Done()
		gen(A, 10)
	}()

	go func() {
		defer wg.Done()
		filter(A, B)
	}()

	go func() {
		defer wg.Done()
		print(B)
	}()
	wg.Wait()
}
