// Есть горутина, которая генерирует случайные числа от 1 до N и отправляет в канал A.
// Вторая горутина читает из A и отправляет в канал B только чётные числа.
// Третья горутина читает из B и выводит числа.
// Можно с WaitGroup, можно с <-time.After()

package main

import (
	"math/rand"
	"fmt"
	"sync"
)

func producer(a chan<- int, n int, wg *sync.WaitGroup) {
	defer wg.Done()
	for {
  		a <- rand.Intn(n + 1)
  	}
}

func reciever(a <-chan int, b chan<- int, wg *sync.WaitGroup) {
	defer wg.Done()
	for {
		n := <-a
		if n % 2 == 0 {
			b <- n
		}
  	}
}

func printer(b <-chan int, wg *sync.WaitGroup) {
	defer wg.Done()
	for {
		n := <-b
		fmt.Println(n)
  	}
}

func do2() {
 	a, b := make(chan int), make(chan int)
  	defer close(a)
  	defer close(b)
  	var wg sync.WaitGroup
  	wg.Add(3)
	
	go producer(a, 1000, &wg)
	go reciever(a, b, &wg)
	go printer(b, &wg)
	
	wg.Wait()
}
