package main

import (
	"errors"
	"time"
)

// Написать функцию after() — аналог time.After().

// after возвращает канал, в котором появится значение через промежуток времени dur.
func after(dur time.Duration) <-chan time.Time {
	// ..
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
