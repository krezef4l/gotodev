// Есть горутина, которая генерирует случайные числа от 1 до N и отправляет в канал A.
// Вторая горутина читает из A и отправляет в канал B только чётные числа.
// Третья горутина читает из B и выводит числа.
// Можно с WaitGroup, можно с <-time.After()
package main

import (
	"fmt"
	"math/rand"
	"sync"
	"time"
)

func numGen(wg *sync.WaitGroup, n int, output chan int) {

	fmt.Printf("Generated nums - ")
	for i := 0; i < n; i++ {
		num := rand.Intn(n)
		output <- num
		fmt.Printf("%v ", num)
	}
	fmt.Printf("\n")

	close(output)
	wg.Done()
}

func numFilter(wg *sync.WaitGroup, input <-chan int, output chan<- int) {
  
	for num := range input {
		if num%2 == 0 {
			output <- num
		}
	}
	close(output)

	wg.Done()
}

func numPrint(wg *sync.WaitGroup, filtInput <-chan int) {
	time.Sleep(10 * time.Millisecond)
	fmt.Printf("Filtered nums - ")
	for num := range filtInput {
		fmt.Printf("%v ", num)
	}
	fmt.Printf("\n")
	wg.Done()
}

func main() {

	var wg sync.WaitGroup

	ch1 := make(chan int, 10)
	ch2 := make(chan int, 10)

	wg.Add(3)
	go numGen(&wg, 20, ch1)
	go numFilter(&wg, ch1, ch2)
	go numPrint(&wg, ch2)

	wg.Wait()
}
