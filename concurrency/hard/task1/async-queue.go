// У нас есть поток задач которые нужно обрабатывать.
// Напишите библиотеку, которая будет принимать на вход задачи и асинхронно обрабатывать их.
// Работать обработчик должен в оперативной памяти.
// В один момент времени выполняется не более N задач,
// и не более М задач могут быть поставлены в очередь.
// Если в очереди нет места, возвращаем ошибку.

package main

import (
	"fmt"
	"math/rand"
	"time"

	taskqueue "github.com/f0xg0sasha/gotodev/tree/main/concurrency/hard/task1/asyncqueue"
)

func main() {
	// Создаем очередь: 2 параллельных задачи, буфер на 6 задачи
	q := taskqueue.New(3, 6)
	q.Run() // Запускаем обработчики

	// Функция задачи
	createTask := func(id int, delay time.Duration) taskqueue.Task {
		return func() {
			fmt.Printf("Task %d started\n", id)
			time.Sleep(delay)
			fmt.Printf("Task %d completed: %d \n", id, id*id)
		}
	}

	// Добавляем задачи
	for i := 1; i <= 15; i++ {
		err := q.AddTask(createTask(i, time.Duration(rand.Intn(3))*500*time.Millisecond))
		if err != nil {
			fmt.Printf("Error adding task %d: %v\n", i, err)
		} else {
			fmt.Printf("Task %d added to queue\n", i)
		}
	}

	// Ждем завершения всех задач
	q.Wait()
	fmt.Println("All tasks completed")
}

// $ go run async-queue.go

// Task 1 added to queue
// Task 2 added to queue
// Task 3 added to queue
// Task 4 added to queue
// Task 5 added to queue
// Task 6 added to queue
// Task 7 added to queue
// Task 8 added to queue
// Task 9 added to queue
// Task 1 started
// Task 2 started
// Task 2 completed: 4
// Task 4 started
// Task 4 completed: 16
// Task 5 started
// Task 3 started
// Error adding task 10: queue is full
// Task 11 added to queue
// Task 12 added to queue
// Error adding task 13: queue is full
// Error adding task 14: queue is full
// Error adding task 15: queue is full
// Task 5 completed: 25
// Task 6 started
// Task 6 completed: 36
// Task 7 started
// Task 1 completed: 1
// Task 8 started
// Task 3 completed: 9
// Task 9 started
// Task 9 completed: 81
// Task 11 started
// Task 7 completed: 49
// Task 12 started
// Task 8 completed: 64
// Task 11 completed: 121
// Task 12 completed: 144
// All tasks completed
