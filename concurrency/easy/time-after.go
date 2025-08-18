package main

import (
	"errors"
	"fmt"
	"time"
)

// Написать функцию after() — аналог time.After():

// возвращает канал, в котором появится значение
// через промежуток времени dur

func after(dur time.Duration) <-chan time.Time {
	ch := make(chan time.Time, 1)
	go func() {
		defer close(ch)

		time.Sleep(dur)
		ch <- time.Time{}
	}()
	return ch
}

func withTimeout(fn func() int, timeout time.Duration) (int, error) {
	var result int

	done := make(chan struct{})
	go func() {
		result = fn()
		close(done)
	}()

	select {
	case <-done:
		return result, nil
	case <-after(timeout): // тут мог быть `<-time.After()`
		return 0, errors.New("timeout")
	}
}

func main() {
	// 1 <nil> (успевает выполниться работа)
	fmt.Println(withTimeout(
		func() int {
			time.Sleep(time.Second * 1) // работает 1 сек
			return 1
		},
		time.Second*1+time.Millisecond*200, // таймаут 1.2 сек
	))

	// 0 timeout (не успевает выполниться работа)
	fmt.Println(withTimeout(
		func() int {
			time.Sleep(time.Second*1 + time.Millisecond*100) // работает 1.1 сек
			return 1
		},
		time.Second*1, // таймаут 1 сек
	))
}
