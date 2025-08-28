// Если тема контекста еще не изучена - пропустить задачу
package main

import (
	"errors"
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

func predictableFunc(timeout time.Duration) (int64, error) {
	var result int64
	done := make(chan struct{})

	go func() {
		defer close(done)
		result = unpredictableFunc()
	}()

	select {
	case <-done:
		return result, nil
	case <-time.After(timeout):
		return 0, errors.New("ERROR: timeout")
	}

}

func main() {
	timeout := time.Duration(1 * time.Second)

	start := time.Now()

	res, err := predictableFunc(timeout)
	if err != nil {
		fmt.Println(err)
		return
	}

	elapsed := float64(time.Since(start)) / 1_000_000

	fmt.Printf("Took %.0f ms\n", elapsed)
	fmt.Println(res)
}
