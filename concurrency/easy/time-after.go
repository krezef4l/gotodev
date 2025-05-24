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
	res := make(chan time.Time)
	go func() {
		time.Sleep(dur)
		res <- time.Now()
	}()
	return res
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
	withTimeout(func() int {
		time.Sleep(time.Second * 2)
		fmt.Println("done")
		return 1
	}, time.Second * 3)
}