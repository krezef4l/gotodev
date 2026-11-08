package main

import (
	"fmt"
	"sync"
	"time"
)

// Написать функцию merge(), которая объединяет произвольное число каналов
// и возвращает смерженный канал. Заранее неизвестно, сколько данных
// придет в каждом из каналов.

func generateInRange(start, stop int) <-chan int {
	out := make(chan int)
	go func() {
		defer close(out)
		for i := start; i < stop; i++ {
			time.Sleep(50 * time.Millisecond)
			out <- i
		}
	}()
	return out
}

func merge(channels ...<-chan int) <-chan int {
	//TODO
	var wg sync.WaitGroup
	out := make(chan int)

	for _, ch := range channels {
		wg.Add(1)
		go func(c <-chan int) {
			defer wg.Done()
			for v := range c {
				out <- v
			}
		}(ch)
	}

	go func() {
		wg.Wait()
		close(out)
	}()
	return out
}

func main() {
	in1 := generateInRange(100, 120)
	in2 := generateInRange(110, 130)

	start := time.Now()
	merged := merge(in1, in2)
	for val := range merged {
		fmt.Print(val, " ")
	}

	fmt.Printf("\nTook %d\n", time.Since(start))
}
