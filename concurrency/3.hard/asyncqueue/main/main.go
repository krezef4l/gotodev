package main

import (
	asyncQlib "asyncQueue/AQLib"
	"fmt"
)

func main() {

	fmt.Println("[main] started")

	//тестовые джобы
	jobs := []asyncQlib.Job{
		func(in, out chan interface{}) {
			out <- 1
			fmt.Println("[job] Done")
			close(out)
		},
		func(in, out chan interface{}) {
			out <- 2
			fmt.Println("[job] Done")
			close(out)
		},
		func(in, out chan interface{}) {
			out <- 3
			fmt.Println("[job] Done")
			close(out)
		},
		func(in, out chan interface{}) {
			out <- 4
			fmt.Println("[job] Done")
			close(out)
		},
		func(in, out chan interface{}) {
			out <- 5
			fmt.Println("[job] Done")
			close(out)
		},
	}

	TS := asyncQlib.TaskSolver{}

	fmt.Println("[main] set limmits")
	TS.SetLimits(6, 2)

	fmt.Println("[main] append the queue")
	for _, curjob := range jobs {
		TS.AppendQueue(curjob)
	}

	fmt.Println("[main] running workers")
	TS.RunWorkers()

	resultValues := TS.GetAllResults()

	//для теста вывожу значения и конвертирую из int, конечно можно заморочиться
	//и собирать все результаты в какойнибудь JSON и потом его анмаршалить,
	//но лаба вроде не про это)
	for _, val := range resultValues {
		fmt.Println(val.(int))
	}
}
