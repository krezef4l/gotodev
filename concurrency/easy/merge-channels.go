// Написать функцию merge(), которая объединяет произвольное число каналов
// и возвращает смерженный канал. Заранее неизвестно, сколько данных
// придет в каждом из каналов.

package main

import (
	"fmt"
	"time"
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
	//TODO
	ar wg sync.WaitGroup

	merged := make(chan int)

	wg.Add(len(channels))

	output := func(gay <-chan int) {
		for nig := range gay {
			merged <- nig
		}
		wg.Done()
	}

	for _, outChan := range channels {
		go output(outChan)
	}

	go func() {
		wg.Wait()
		close(merged)
	}()

	return merged
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
