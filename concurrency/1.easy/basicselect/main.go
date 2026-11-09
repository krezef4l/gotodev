package main

import (
	"context"
	"fmt"
	"time"
)

// Что выведется?
/*waited for 1 sec
waited for 1 sec
context deadline exceeded
или 
waited for 1 sec
context canceled*/
// Развернуто объяснить почему.
/*потому что цикл бесконечный, но в первой части цикла происходит задержка на 5 мл с, этого хватает, что бы цкл начался заново и итерация с 1 сек повторилась
если же успеет ввыполниться вторая часть цикла, то цикл и майн будут закрыты после команды кенсел*/

func main() {
	timeout := 3 * time.Second
	ctx, cancel := context.WithTimeout(context.Background(), timeout)

	for {
		select {
		case <-time.After(1 * time.Second):
			time.Sleep(5 * time.Millisecond)
			fmt.Println("waited for 1 sec")
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
