// Реализовать пул из 3 воркеров, которые:
// - получают задачи (в задачах спим и что-то печатаем, например) из общего канала.
// - вычисляют квадрат числа и отправляют результат в общий канал.
// Главная горутина создаёт N задач, распределяет их по воркерам и выводит результаты.

package main

import (
	"fmt"
	"sync"
	"time"
)

func Worker(tasks <-chan int, results chan<- int, id int, wg *sync.WaitGroup) {
	for num := range tasks {
		time.Sleep(time.Millisecond)
		fmt.Printf("[worker %v] Sending result by worker %v\n", id, id)
		results <- num * num
	}
	wg.Done()
}

func main() {
	var wg sync.WaitGroup
	bufferSize := 10
	tasks := make(chan int, bufferSize)
	results := make(chan int, bufferSize)

	for i := range 3 {
		wg.Add(1)
		go Worker(tasks, results, i, &wg)
	}
	fmt.Printf("[main] Wrote %v tasks\n", bufferSize)

	for i := range bufferSize {
		tasks <- i
	}
	close(tasks)

	wg.Wait()
	close(results)
	var counter int
	for num := range results {
		counter++
		fmt.Println("[main] Result", counter, ":", num)
	}
	fmt.Println("[main] main() stopped")
}