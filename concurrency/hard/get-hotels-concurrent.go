package main

import (
	"fmt"
	"sync"
	"time"
)

// Есть поток данных, в виде идентификаторов отелей (hotelIDs), для каждого отеля нужно:
// 1) сделать поисковый запрос (search), запрос выполняется 500ms;
// 2) отправить результаты в другой канал;
// 3) прочитать результаты из канала и вывести на экран.

type SearchResult struct {
	HotelID int
}

func main() {
	hotelIDs := getHotels()

	// Код здесь. Остальные функции и их сигнатуры менять нельзя. Допускается использовать функции-обертки.

	// Создаю результирующий в который мы будем писать и из него же читать
	resultCh := make(chan SearchResult)
	// Чтоб закрыть канал после завершения работы горутин
	wg := &sync.WaitGroup{}

	// Читаем канал
	for id := range hotelIDs {
		wg.Add(1)

		// Ищем инфу и пишем в канал
		go func(id int) {
			resultCh <- search(id)
			wg.Done()
		}(id)
	}
	// Закрываем канал
	go func() {
		wg.Wait()
		close(resultCh)
	}()

	// Читаем результирующий канал
	for i := range resultCh {
		fmt.Println(i)
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
