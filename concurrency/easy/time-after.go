package main

import (
	"errors"
	"time"
)

// Написать функцию after() — аналог time.After():

// возвращает канал, в котором появится значение
// через промежуток времени dur
func after(dur time.Duration) <-chan time.Time {
	done := make(chan time.Time, 1)

	go func() {
		defer close(done)
		time.Sleep(dur)
		done <- time.Now()
	}()

	return done
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
