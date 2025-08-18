// Написать функцию merge(), которая объединяет произвольное число каналов
// и возвращает смерженный канал. Заранее неизвестно, сколько данных
// придет в каждом из каналов.

package main

import (
	"fmt"
	"sync"
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

	// вейтгрупа для того, чтобы успеть считать все данные из каналов, а после закрыть нащ результирующий
	wg := &sync.WaitGroup{}

	// создаем и инициализируем наш результирующий канал
	resCh := make(chan int)

	// итерируемся по слайсу каналов (не уверен, что по слайсу, но поидее так и должно быть)
	for _, ch := range channels {
		// инкриментим счетчик
		wg.Add(1)
		go func(ch <-chan int) {
			// читаем из канала
			for item := range ch {
				// пишем в наш результирующий канал канал
				resCh <- item
			}
			// декриментим счетчик
			wg.Done()
		}(ch)
	}

	// небольшая махинация для того, чтоб:
	// 1. не закрыть канал слишком рано, а потом горутиной писать в закрытый канал -> panic(),
	// 2. вообще закрыть в целом, чтобы не словить deadlock, потому что "for val := range merged" будет ждать, а мы уже не пишем -> дедлок
	go func() {
		wg.Wait()
		close(resCh)
	}()

	return resCh
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
