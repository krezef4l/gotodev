package main

import (
	"context"
	"errors"
	"sync"
)

// У нас есть поток задач которые нужно обрабатывать.
// Напишите библиотеку, которая будет принимать на вход задачи и асинхронно обрабатывать их.
// Работать обработчик должен в оперативной памяти.
// В один момент времени выполняется не более N задач,
// и не более М задач могут быть поставлены в очередь.
// Если в очереди нет места, возвращаем ошибку.

type Queue[T any] interface {
	Add(value T) error
	Get() <-chan T
}


type Job struct {
	Run func()
}

func NewJob(runnable func()) *Job {
	return &Job{
		Run: runnable,
	}
}

type JobQueue struct {
	queue chan *Job
}

func (jq *JobQueue) Add(job *Job) error {
	if len(jq.queue) == cap(jq.queue) {
		return errors.New("queue is full")
	}

	jq.queue <- job
	return nil
}

func (jq *JobQueue) Get() <-chan *Job {
	return jq.queue
}

type Worker struct {
	jobs chan *Job
	ctx context.Context
	wp *WorkerPool
}

func (w *Worker) Start() {
	go func () {
		for {
			w.wp.pool <- w
			select {
			case job := <-w.jobs:
				w.wp.wg.Add(1)
				job.Run()
				w.wp.wg.Done()
			case <-w.ctx.Done():
				return
			}
		}
	}()
}

func NewWorker(ctx context.Context, wp *WorkerPool) *Worker {
	return &Worker{
		jobs: make(chan *Job),
		ctx: ctx,
		wp: wp,
	}
}

func NewJobQueue(size int) *JobQueue {
	return &JobQueue{
		queue: make(chan *Job, size),
	}
}

type WorkerPool struct {
	queue Queue[*Job]
	pool chan *Worker
	wg sync.WaitGroup
}

func NewWorkerPool(jobQueueSize int, numberOfWorkers int) *WorkerPool {
	return &WorkerPool{
		queue: NewJobQueue(jobQueueSize),
		pool: make(chan *Worker, numberOfWorkers),
		wg: sync.WaitGroup{},
	}
}

func (wp *WorkerPool) Start(ctx context.Context) {
	println("Starting worker pool")
	for range len(wp.pool) {
		worker := NewWorker(ctx, wp)
		worker.Start()
	}

	go func() {
		for {
			select {
				case job := <-wp.queue.Get():
					worker := <-wp.pool
					worker.jobs <- job
				case <-ctx.Done():
					wp.wg.Wait()
					println("Shutting down worker pool")
					return
			}
		}
	}()
	
	println("Worker pool started")
}

func (wp *WorkerPool) AddToQueue(job *Job) error {
	return wp.queue.Add(job)
}