package main

import (
	"context"
	"fmt"
	"time"
)

// Что выведется?
// Развернуто объяснить почему.

func main() {
	//создали контекст по таймауту равным 3 секунды
	timeout := 3 * time.Second
	ctx, cancel := context.WithTimeout(context.Background(), timeout)

	for {
		select {
		/*
			time.After(). Эта библиотечная функция возвращает канал,
			который изначально пуст, а через timeout времени отправляет в него значение.
			Благодаря этому select выберет нужный вариант, а именно
			'case <-time.After(1 * time.Second)' дважды
		*/
		case <-time.After(1 * time.Second):
			time.Sleep(5 * time.Millisecond)
			fmt.Println("waited for 1 sec")
			/*
				не выполняем 'case <-time.After(2 * time.Second)' и 'case <-time.After(3 * time.Second)'
				по причине наступления таймаута контекста
				сигнал об отмене приходит из канала ctx.Done()
			*/
		case <-time.After(2 * time.Second):
			fmt.Println("waited for 2 sec")
			cancel()
		case <-time.After(3 * time.Second):
			fmt.Println("waited for 3 sec")
		case <-ctx.Done():
			fmt.Println(ctx.Err())
			return
		}
	}
}

/* Ответ:
waited for 1 sec
waited for 1 sec
context deadline exceeded
*/
