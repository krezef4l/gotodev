package main

import (
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
// (например, 1 секунду).]
//
// Если "длинная" функция отработала за это время - отлично, возвращаем
// результат. Если нет - возвращаем ошибку. Результат работы в этом случае нам
// неважен.
//
// Дополнительно нужно измерить, сколько выполнялась эта функция (просто вывести
// в лог). Сигнатуру функцию обёртки (predictableFunc) менять можно.

func predictableFunc() int64 {
	return unpredictableFunc()
}

func main() {
	res := predictableFunc()
	fmt.Println(res)
}
