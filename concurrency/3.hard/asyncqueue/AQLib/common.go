package asyncQlib

import (
	"fmt"
)

// приходящие ф-ции джобы
type Job func(in, out chan interface{})

func worker(idx int, inCh <-chan interface{}, outCh chan<- interface{}) {

	fmt.Printf("[worker] Worker %d started\n", idx)

	for {
		select {
		case <-CancelSignal:
			fmt.Printf("[worker] Worker %d killed\n", idx)
			return

		case job, ok := <-inCh:
			if !ok {
				// no more jobs
				return
			}

			// Run the job
			curJob := job.(Job)

			jobInCh := make(chan interface{}, 1)
			jobOutCh := make(chan interface{}, 1)

			go curJob(jobInCh, jobOutCh)

			// If your job doesn't need an input, skip this
			// jobInCh <- someInput
			close(jobInCh)

			// read result
			result := <-jobOutCh

			// send result back
			outCh <- result
		}
	}
}

type QueueOverflowError struct {
	Err error
}

func (e *QueueOverflowError) Error() string {
	return "queue overrflow" + e.Err.Error()
}

type KillWorkersError struct {
	Err error
}

func (e *KillWorkersError) Error() string {
	return "not all workers killed" + e.Err.Error()
}

/*
	for job := range inCh {
		jobInCh := make(chan interface{})
		jobOutCh := make(chan interface{})
		curJob := job.(Job)
		curJob(jobInCh, jobOutCh)
		close(jobOutCh)

		outCh = jobOutCh
	}
*/

/*
for job := range inCh {
		select {
		case <-CancelSignal:
			break
		default:
			jobInCh := make(chan interface{})
			jobOutCh := make(chan interface{})
			curJob := job.(Job)
			curJob(jobInCh, jobOutCh)
			close(jobOutCh)

			outCh = jobOutCh
		}
	}
*/
