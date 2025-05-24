// Написать функцию merge(), которая объединяет произвольное число каналов
// и возвращает смерженный канал. Заранее неизвестно, сколько данных
// придет в каждом из каналов.

package main

import (
	"fmt"
	"time"
	"sync"
)

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
	res := make(chan int, len(channels))
	var wg sync.WaitGroup
	wg.Add(len(channels))
	
	for _, ch := range channels {
		go func(ch <-chan int) {
			for x := range ch {
				res <- x
			}
			wg.Done()
		}(ch)
	}
	
	go func() {
		wg.Wait()
		close(res)
	}()
	
	return res
}

func do4() {
	in1 := generateInRange(100, 120)
	in2 := generateInRange(110, 130)

	start := time.Now()
	merged := merge(in1, in2)
	for val := range merged {
		fmt.Print(val, " ")
	}

	fmt.Printf("\nTook %d\n", time.Since(start))
}
