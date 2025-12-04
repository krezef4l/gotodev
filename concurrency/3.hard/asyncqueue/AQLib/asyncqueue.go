// У нас есть поток задач, которые нужно обрабатывать. Напишите библиотеку,
// которая будет принимать на вход задачи и асинхронно обрабатывать их.
//
// 1) Работать обработчик должен в оперативной памяти.
// 2) В один момент времени выполняется не более N задач, и не более М задач
// могут быть поставлены в очередь.
// 3) Если в очереди нет места, возвращаем ошибку.
package asyncQlib

import (
	"fmt"
	"log"
	"sync"
)

var (
	//счётчик горутин
	GoroutineCount int
	//канал отмены
	CancelSignal chan interface{}
)

type TaskSolver struct {
	wg sync.WaitGroup
	//лимит одновременно выполняющихся задач (грубо говоря макс. кол-во горутин для выполнения задач)
	TaskLimit int
	//лимит задач в очереди
	QueueLimit int
	JobQueue   (chan Job)
	Result     []interface{}
	queueErr   *QueueOverflowError
	workersErr *KillWorkersError
}

// Set limits (queue_size, task_at_the_same_time)
func (t *TaskSolver) SetLimits(m, n int) {
	t.TaskLimit = n - 1
	t.QueueLimit = m - 1
	t.JobQueue = make(chan Job, t.QueueLimit)
	log.Println("[TS] Hell Yeah, you set the queue and tasks limits!")
}

func (t *TaskSolver) AppendQueue(job Job) error {

	if len(t.JobQueue) >= t.QueueLimit {
		return t.queueErr
	}
	t.JobQueue <- job
	fmt.Println("[TS] Task succssfully appended to the queue")
	return nil
}

func (t *TaskSolver) RunWorkers() {

	close(t.JobQueue)
	fmt.Println("[TS] try start all workers")

	inputJob := make(chan interface{})
	result := make(chan interface{})

	// Start workers
	for i := 0; i < t.TaskLimit; i++ {
		t.wg.Add(1)
		GoroutineCount++
		go func(id int) {
			defer t.wg.Done()
			worker(id, inputJob, result)
		}(i)
	}

	// Send jobs to workers
	go func() {
		for job := range t.JobQueue {
			inputJob <- job
		}
		close(inputJob) // IMPORTANT
	}()

	fmt.Println("[TS] all jobs ready to be solved")

	fmt.Println("[TS] all workers started")

	// Close result only when workers finish
	go func() {
		t.wg.Wait()
		close(result) // IMPORTANT
	}()

	fmt.Println("[TS] write results")
	// Collect results
	for v := range result {
		t.Result = append(t.Result, v)
	}

	fmt.Println("Results done")
}

func (t *TaskSolver) GetAllResults() []interface{} {
	return t.Result
}

func (t *TaskSolver) KillWorkers() error {

	CancelSignal = make(chan interface{})
	defer close(CancelSignal)
	cancel := make([]interface{}, t.TaskLimit)
	fmt.Println("[TS] start killing workers")
	for i := range cancel {
		cancel[i] = struct{}{}
	}
	CancelSignal <- cancel

	if GoroutineCount != 0 {
		return t.workersErr
	}
	return nil
}
