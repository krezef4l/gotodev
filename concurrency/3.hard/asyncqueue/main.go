package main

// У нас есть поток задач, которые нужно обрабатывать. Напишите библиотеку,
// которая будет принимать на вход задачи и асинхронно обрабатывать их.
//
// 1) Работать обработчик должен в оперативной памяти.
// 2) В один момент времени выполняется не более N задач, и не более М задач
// могут быть поставлены в очередь.
// 3) Если в очереди нет места, возвращаем ошибку.

import (
	"errors"
	"fmt"
	"sync"
	"time"
)

var ErrQueueFull = errors.New("очередь заполнена")

type Pool struct {
	queue chan func()
	wg    sync.WaitGroup
}

// New workers = N (сколько задач одновременно),
// queueSize = M (сколько задач ждут в очереди).
func New(workers, queueSize int) *Pool {
	p := &Pool{
		queue: make(chan func(), queueSize),
	}
	p.wg.Add(workers)
	for range workers {
		go p.worker()
	}
	return p
}

// Submit если в очереди нет места сразу ошибка
func (p *Pool) Submit(task func()) error {
	select {
	case p.queue <- task:
		return nil
	default:
		return ErrQueueFull
	}
}

func (p *Pool) worker() {
	defer p.wg.Done()
	for task := range p.queue {
		task()
	}
}

func (p *Pool) Close() {
	close(p.queue)
	p.wg.Wait()
}

func main() {
	pool := New(4, 10)
	defer pool.Close()

	for i := range 20 {
		err := pool.Submit(func() {
			time.Sleep(100 * time.Millisecond)
			fmt.Println("task", i)
		})
		if err != nil {
			fmt.Println("task", i, "->", err)
		}
	}
}
