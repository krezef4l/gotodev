package main

import (
	"context"
	"fmt"
	"time"
)

// Что выведет?

/*
 На каждой итерации будет создаваться новый time.After во всех кейсах и заново течь время,
 таким образом мы будем все время заходить только в первый кейс так как он быстрее всего исполняется.
 Через три секунды откроется последний кейс так как истечет таймаут и мы сразу туда зайдем, так как вначале минимум
 одну секунду во всех каналах не будет значений.
*/ 

/*
На третьей итерации сначала будет доступен канал ctx.Done() через суммарное время 3s, а потом будет 
доступен time.After() через суммарное время 3.10. Это не гарантирует, что выберется ctx.Done
так как планировщик может задержаться какое-то время и 0.10 слишком незначительное время.
*/

func do1() {
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


