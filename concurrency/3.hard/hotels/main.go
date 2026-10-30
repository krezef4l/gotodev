package main

import (
	"fmt"
	"sync"
	"time"
)

// Есть поток данных в виде идентификаторов отелей (hotelIDs), для каждого
// отеля нужно:
// 1) сделать поисковый запрос (search), запрос выполняется 500ms;
// 2) отправить результаты в другой канал;
// 3) прочитать результаты из канала и вывести на экран.

type SearchResult struct {
	HotelID int
}

func main() {
	hotelIDs := getHotels()
	results := make(chan SearchResult)

	// Код здесь. Остальные функции и их сигнатуры менять нельзя. Допускается
	// использовать функции-обертки.
	var wg sync.WaitGroup
	for id := range hotelIDs {
		wg.Add(1)
		go func() {
			defer wg.Done()
			results <- search(id)
		}()
	}

	go func() {
		wg.Wait()
		close(results)
	}()

	for i := range results {
		fmt.Println(i.HotelID)
	}
}

func search(hotelID int) SearchResult {
	time.Sleep(time.Millisecond * 500)
	return SearchResult{HotelID: hotelID}
}

func getHotels() chan int {
	ch := make(chan int)
	go func() {
		for i := 1; i <= 10; i++ {
			ch <- i
		}
		close(ch)
	}()
	return ch
}
