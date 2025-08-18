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

var tasksCount = 20
var workersCount = 3

func worker(id int, nums <-chan int, results chan<- []int, wg *sync.WaitGroup) {
	defer wg.Done()
	for num := range nums {
		time.Sleep(time.Millisecond * 500) //имитация хоть небольшой нагрузки, чтоб видно было, что все воркеры работают, когда нагрузки нет больше 90% задач берет один воркер.
		results <- []int{id, num * num}
	}
}

func main() {
	tasks := make(chan int, tasksCount)
	results := make(chan []int, tasksCount)

	wg := new(sync.WaitGroup)

	// здесь я создал 3 воркера, добавил waitgroup чтобы дождаться завершения всех воркеров
	for i := 1; i <= workersCount; i++ {
		wg.Add(1)
		go worker(i, tasks, results, wg)
	}

	// тут я с помощью горутины посылаю числа в канал "tasks"
	go func() {
		for i := 0; i < tasksCount; i++ {
			tasks <- i
		}
		// хорошее правило закрывать канал там же где и создали
		close(tasks)
	}()

	// тут я после завершения работы воркеров - закрываю канал
	go func() {
		wg.Wait()

		// хорошее правило закрывать канал там же где и создали
		close(results)
	}()

	// тут я читаю из канала результат
	for res := range results {
		fmt.Printf("Worker №%d result: %d\n", res[0], res[1])
	}

}
