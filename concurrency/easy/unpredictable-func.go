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

func predictableFunc() (int64, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 1*time.Second)
	defer cancel()
	ch := make(chan int64, 1)
	var res int64
	go func() {
		ch <- unpredictableFunc()
	}()
	select {
	case <-ctx.Done():
		return 0, ctx.Err()
	case res = <-ch:
		return res, nil
	}
}

func main() {
	start := time.Now()
	res, err := predictableFunc()
	elapsed := time.Since(start)
	fmt.Println(res)
	fmt.Println(err)
	fmt.Println("Время затраченное: ", elapsed)
}
