package main

import (
	"context"
	"fmt"
	"time"
)

// Что выведет?

/*
waited for 1 sec
waited for 1 sec
context deadline exceeded

*/

func main() {
	timeout := 3 * time.Second
	// Создается контекст с дедлайном в 3 секунды
	ctx, cancel := context.WithTimeout(context.Background(), timeout)

	for {
		select {
		case <-time.After(1 * time.Second):
			// В цикле мы попадаем в 1 кейс после истечения 1 секунды.
			// После спим 5 мс
			// До истечения контекста с учетом паузы мы успеваем попасть в этот кейс 2 раза
			time.Sleep(5 * time.Millisecond)
			fmt.Println("waited for 1 sec")
		case <-time.After(2 * time.Second):
			fmt.Println("waited for 2 sec")
			cancel()
		case <-time.After(3 * time.Second):
			fmt.Println("waited for 3 sec")
		case <-ctx.Done():
			// При наступлении 3 секунд контекст завершается с ошибкой
			fmt.Println(ctx.Err())
			return
		}
	}
}
