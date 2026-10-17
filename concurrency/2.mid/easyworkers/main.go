package main

import (
	"fmt"
	"sync"
	"time"
)

// worker получает задачу, ждет и возводит значение в квадрат
func worker(id int, task <-chan int, res chan<- int, wg *sync.WaitGroup) {
	defer wg.Done()

	for num := range task {
		fmt.Printf("Worker %d: task %d\n", id, num)
		//работа
		time.Sleep(100 * time.Millisecond)
		square := num * num
		fmt.Printf("Worker %d: %d^2 = %d\n", id, num, square)
		res <- square
	}
}

func main() {
	const nWorker = 3
	const nTask = 10

	res := make(chan int, nTask)
	tasks := make(chan int, nTask)

	var wg sync.WaitGroup

	for i := 0; i < nWorker; i++ {
		wg.Add(1)
		go worker(i+1, tasks, res, &wg)
	}

	for i := 1; i <= nTask; i++ {
		tasks <- i
	}
	close(tasks)

	go func() {
		wg.Wait()
		close(res)
	}()

	summ := 0
	for results := range res {
		fmt.Printf("Result: %d\n", results)
		summ += results
	}

	fmt.Printf("Sum square: %d\n", summ)

}

// Реализовать пул из 3-х воркеров, которые:
// - получают задачи (в задачах просто спим и что-то печатаем, например) из
// общего канала;
// - вычисляют квадрат числа и отправляют результат в общий канал.
// Главная горутина создаёт N задач, распределяет их по воркерам и выводит
// результаты.
