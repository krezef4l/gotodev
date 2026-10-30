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

// Есть функция, работающая неопределённо долго и возвращающая число. Её тело
// нельзя изменять (представим, что внутри сетевой запрос).
func unpredictableFunc() int64 {
	rnd := rand.Int63n(5000)
	time.Sleep(time.Duration(rnd) * time.Millisecond)
	return rnd
}

// Нужно изменить функцию-обёртку, которая будет работать с заданным таймаутом
// (например, 1 секунду). Если "длинная" функция отработала за это время -
// отлично, возвращаем результат. Если нет - возвращаем ошибку. Результат работы
// в этом случае нам не важен.
//
// Дополнительно нужно измерить, сколько выполнялась эта функция (просто вывести
// в лог). Сигнатуру функцию обёртки менять можно.

func predictableFunc() (int64, error) {

	resCh := make(chan int64)
	ctx, cancel := context.WithTimeout(context.Background(), 2*time.Second)
	defer cancel()

	start := time.Now()

	go func() {
		resCh <- unpredictableFunc()
	}()

	select {
	case timeRand := <-resCh:
		fmt.Println("Time:", time.Since(start))
		return timeRand, nil
	case <-ctx.Done():
		fmt.Println("Time:", time.Since(start))
		return 0, ctx.Err()
	}
}

func main() {
	res, err := predictableFunc()
	if err != nil {
		fmt.Println("Функция не выполнилась..", err)
	} else {
		fmt.Println("Функция  выполнилась", res)
	}
}

//Выполнил через контекст. Так проще установить таймер
//и поставил счетчик для вычисления + в конце пишем выполнено или нет.
