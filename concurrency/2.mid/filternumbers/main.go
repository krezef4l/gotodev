package main


import (
	"fmt"
	"math/rand"
	"sync"
)

func generator(n int, out chan<- int) {
	defer close(out)

	for i := 0; i < n; i++ {
		num := rand.Intn(n) + 1
		out <- num
	}
}

func filter(in <-chan int, out chan<- int) {
	defer close(out)

	for num := range in {
		if num%2 == 0 {
			out <- num
		}
	}
}

func printer(in <-chan int) {
	for num := range in {
		fmt.Println(num)
	}
}

func main() {
	n := 10

	a := make(chan int)
	b := make(chan int)

	var wg sync.WaitGroup
	wg.Add(3)

	go func() {
		defer wg.Done()
		generator(n, a)
	}()

	go func() {
		defer wg.Done()
		filter(a, b)
	}()

	go func() {
		defer wg.Done()
		printer(b)
	}()

	wg.Wait()
}
// Есть горутина, которая генерирует случайные числа от 1 до N и отправляет в
// канал A.
// Вторая горутина читает из A и отправляет в канал B только чётные числа.
// Третья горутина читает из B и выводит числа.
// Можно с WaitGroup, можно с <-time.After().
