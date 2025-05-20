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
	// ..
	timeChan := make(chan time.Time, 1)
	// go func() {
	time.Sleep(dur)
	timeChan <- time.Now()
	close(timeChan)
	// }()
	return timeChan
}

func withTimeout(fn func() int, timeout time.Duration) (int, error) {
	var result int

	done := make(chan struct{})
	go func() {
		result = fn()
		time.Sleep(5 * time.Second)
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
	getInt := func() int {
		return 8
	}
	time.Sleep(time.Millisecond)
	timeOut := 1000 * time.Millisecond
	result, err := withTimeout(getInt, timeOut)
	fmt.Printf("Test 1: result=%d, err=%v\n", result, err)
}
