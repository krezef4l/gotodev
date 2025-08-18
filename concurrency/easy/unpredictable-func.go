// Если тема контекста еще не изучена - пропустить задачу
package main

import (
	"context"
	"fmt"
	"math/rand"
	"time"
)

func init() {
	rand.Seed(time.Now().UnixNano())
}

// Есть функция, работающая неопределённо долго и возвращающая число.
// Её тело нельзя изменять (представим, что внутри сетевой запрос).
func unpredictableFunc() int64 {
	rnd := rand.Int63n(5000)
	time.Sleep(time.Duration(rnd) * time.Millisecond)
	return rnd
}

// Нужно изменить функцию-обёртку, которая будет работать с заданным таймаутом (например, 1 секунду).
// Если "длинная" функция отработала за это время - отлично, возвращаем результат.
// Если нет - возвращаем ошибку. Результат работы в этом случае нам не важен.
//
// Дополнительно нужно измерить, сколько выполнялась эта функция (просто вывести в лог).
// Сигнатуру функцию обёртки менять можно.

func predictableFunc() int64 {
	start := time.Now()
	var result int64
	done := make(chan struct{})

	ctx, _ := context.WithTimeout(context.Background(), time.Second)

	go func(int64) {
		result = unpredictableFunc()
		close(done)
	}(result)

	select {
	case <-done:
		since := time.Since(start)
		fmt.Println("Work my func:", since, "and my result:", result)
		return result
	case <-ctx.Done():
		since := time.Since(start)
		fmt.Println("Work my func:", since)
		return 0
	}
}

func main() {
	res := predictableFunc()
	fmt.Println(res)
}
