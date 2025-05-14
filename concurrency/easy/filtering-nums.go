// Есть горутина, которая генерирует случайные числа от 1 до N и отправляет в канал A.
// Вторая горутина читает из A и отправляет в канал B только чётные числа.
// Третья горутина читает из B и выводит числа.
// Можно с WaitGroup, можно с <-time.After()

package main

import (
	"fmt"
	"math/rand"
	"sync"
)

func randomGen(wg *sync.WaitGroup, max int, A chan<- int) {
	randomInt := rand.Intn(max) + 1
	A <- randomInt
	fmt.Println("[randomGen] randomGen() recorded num: ", randomInt)
	wg.Done()
}

func evenRec(wg *sync.WaitGroup, A <-chan int, B chan<- int) {
	num := <-A
	if num%2 == 0 {
		B <- num
		fmt.Println("[evenRec] evenRec() recorded even num: ", num)
	}
	wg.Done()
}

func printNums(wg *sync.WaitGroup, B <-chan int) {
	for num := range B {
		fmt.Println("[printNums] printNums() printed even num: ", num)
	}
	wg.Done()
}

func main() {
	fmt.Println("[main] main() started")
	var wgGen sync.WaitGroup
	var wgRec sync.WaitGroup
	var wgPrint sync.WaitGroup
	count := 100
	A := make(chan int)
	B := make(chan int, count/2+10)

	for range count {
		go randomGen(&wgGen, count, A)
		wgGen.Add(1)

		go evenRec(&wgRec, A, B)
		wgRec.Add(1)
	}

	wgGen.Wait()
	close(A)

	wgRec.Wait()
	close(B)

	wgPrint.Add(1)
	go printNums(&wgPrint, B)

	wgPrint.Wait()
	fmt.Println("[main] main() stopped")
}
