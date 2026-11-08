package main

import (
	"errors"
	"fmt"
	"time"
)

// Написать функцию after() — аналог time.After().

// after возвращает канал, в котором появится значение через промежуток времени
// dur.
func after(dur time.Duration) <-chan time.Time {
	out := make(chan time.Time, 1)

	go func() {
		time.Sleep(dur)
		out <- time.Now()

	}()
	return out
}

// Тело функции ниже изменять нельзя.
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
	case <-after(timeout): // тут мог быть `<-time.After()`.
		return 0, errors.New("timeout")
	}
}

func main() {
	start := time.Now()
	res, err := withTimeout(func() int {
		time.Sleep(time.Second * 10)
		fmt.Println("done")
		return 1
	}, time.Second*2)

	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(res)

	fmt.Println(time.Since(start))
}
