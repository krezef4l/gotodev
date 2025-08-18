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

var (
	wg = &sync.WaitGroup{}
	N  = 100
)

func gorutine1(n int, A chan<- int) {
	defer wg.Done()
	defer close(A)

	for i := 0; i < 10; i++ {
		rnd := rand.Intn(n) + 1
		A <- rnd
	}

}

func gorutine2(A <-chan int, B chan<- int) {
	defer wg.Done()
	defer close(B)

	for num := range A {
		if num%2 == 0 {
			B <- num
		}
	}
}

func gorutine3(B <-chan int) {
	defer wg.Done()

	for sorted_nums := range B {
		fmt.Println(sorted_nums)
	}
}

func main() {
	A := make(chan int)
	B := make(chan int)

	wg.Add(3)
	go gorutine1(N, A)
	go gorutine2(A, B)
	go gorutine3(B)

	wg.Wait()
}
