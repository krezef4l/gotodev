// Реализовать пул из 3 воркеров, которые:
// - получают задачи (в задачах спим и что-то печатаем, например) из общего канала.
// - вычисляют квадрат числа и отправляют результат в общий канал.
// Главная горутина создаёт N задач, распределяет их по воркерам и выводит результаты.
package main

import (
	"log"
	"math/rand"
	"sync"
)

// создаю воркера который будет получать таски из общего для всех воркеров канала tasks и отправлятб в общий канал results

func worker(wg *sync.WaitGroup, inputChan <-chan int, output chan<- int, idx int) {

	defer wg.Done()

	log.Printf("[worker] Worker #%v started\n", idx)
	//time.Sleep(200 * time.Millisecond)

	for num := range inputChan {
		log.Printf("worker #%v calculate square of num %v\n", idx, num) // воркер печатает из общего канала число которое он будет возводить в степень
		output <- num * num
	}
}

func main() {

	log.Println("[main] main() started")
	var wg sync.WaitGroup // создаю вейтгруппу чтобы это все не посыпалось
	var tasks = make(chan int, 10)
	var results = make(chan int, 10) // создаю каналы

	// запускаю три воркера
	for i := 0; i < 3; i++ {
		wg.Add(1)
		go worker(&wg, tasks, results, i+1)
	}

	// определяю энное число задач
	N := rand.Intn(10)

	// Закидываю таски в канал
	for i := 1; i <= N; i++ {
		if i%2 != 0 {
			tasks <- i * 2
		} else {
			tasks <- (i * 3) - 3
		}
	}
	log.Printf("[main] Wrote %v tasks\n", N)

	//после отправки в канал всех задач, закрываю канал
	close(tasks)

	//ждём пока все воркеры доработают
	go func() {
		wg.Wait()
		close(results)
	}()

	for i := range results {
		result := <-results
		log.Println("[main] Result", i, ":", result) // выводим результаты
	}
	log.Println("[main] main() stopped")

}


