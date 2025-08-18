package taskqueue

import (
	"errors"
	"sync"
)

// Task представляет функцию для выполнения
type Task func()

// TaskQueue управляет выполнением задач с ограничениями
type TaskQueue struct {
	concurrency int
	maxQueue    int
	tasks       chan Task
	wg          sync.WaitGroup
	mu          sync.Mutex
	activeCount int
}

// New создает новую очередь задач
func New(concurrency, maxQueue int) *TaskQueue {
	return &TaskQueue{
		concurrency: concurrency,
		maxQueue:    maxQueue,
		tasks:       make(chan Task, maxQueue),
	}
}

// Run запускает обработчики задач
func (tq *TaskQueue) Run() {
	for i := 0; i < tq.concurrency; i++ {
		go tq.worker()
	}
}

// Воркер для обработки задач
func (tq *TaskQueue) worker() {
	defer close(tq.tasks)
	for task := range tq.tasks {
		tq.mu.Lock()
		tq.activeCount++
		tq.mu.Unlock()

		// Таска
		task()

		tq.mu.Lock()
		tq.activeCount--
		tq.wg.Done()
		tq.mu.Unlock()
	}
}

// AddTask добавляет задачу в очередь
func (tq *TaskQueue) AddTask(task Task) error {
	if tq.maxQueue > 0 && len(tq.tasks) >= tq.maxQueue {
		return errors.New("queue is full")
	}

	tq.wg.Add(1)
	tq.tasks <- task
	return nil
}

// Wait ожидает завершения всех запущенных задач
func (tq *TaskQueue) Wait() {
	tq.wg.Wait()
}
