// Реализовать пул из 3 воркеров, которые:
// - получают задачи (в задачах спим и что-то печатаем, например) из общего канала.
// - вычисляют квадрат числа и отправляют результат в общий канал.
// Главная горутина создаёт N задач, распределяет их по воркерам и выводит результаты.
package main

import (
	"context"
	"fmt"
	"sync"
	"time"
)

const (
	NumberOfWorkers     = 3
	QueueSize = 3
)
type Job struct {
	Run func()
}

func NewJob(runnable func()) *Job {
	return &Job{
		Run: runnable,
	}
}
type Worker struct {
	jobs chan *Job
	wp *WorkerPool
	ctx context.Context
}

func NewWorker(ctx context.Context, wp *WorkerPool) *Worker {
	return &Worker{
		jobs: make(chan *Job),
		wp: wp, 
		ctx: ctx,
	}
}

func (w *Worker) Start() {
	go func() {
		for {
			w.wp.pool <- w
	
			select {
			case job := <-w.jobs:
				job.Run()
				w.wp.wg.Done()
			case <-w.ctx.Done():
				return
			}
		}
	}()

}

type WorkerPool struct {
	pool chan *Worker
	queue chan *Job
	wg sync.WaitGroup
}

func NewWorkerPool() *WorkerPool {
	return &WorkerPool{
		pool: make(chan *Worker, NumberOfWorkers),
		queue: make(chan *Job, QueueSize),
		wg: sync.WaitGroup{},
	}
}

func (wp *WorkerPool) Add(job *Job) {
	wp.queue <- job
}

func (wp *WorkerPool) Start(ctx context.Context) {
	for i := 0; i < NumberOfWorkers; i++ {
		worker := NewWorker(ctx, wp)
		worker.Start()
	}

	go func() {
			for {
			select {
			case job := <-wp.queue:
				worker := <-wp.pool
				worker.jobs <- job
			case <-ctx.Done():
				return
			}
		}
	}()
}

func main() {
	wp := NewWorkerPool()
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()
	wp.Start(ctx)
	n := 10
	res := make(chan int, n)

	for i := 0; i < n; i++ {
		value := i
		job := NewJob(func() {
			time.Sleep(time.Second)
			fmt.Printf("Calculating %d\n", value)
			res <- value * value
		})
		wp.wg.Add(1)
		wp.Add(job)
	}

	wp.wg.Wait()
}