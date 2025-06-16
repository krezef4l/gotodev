package main

import (
	"errors"
	"time"
)

// Написать функцию after() — аналог time.After():

// возвращает канал, в котором появится значение      // какое значение будет передаваться я тоже не понял нихуя
// через промежуток времени dur
func after(dur time.Duration) <-chan time.Time {
	var timeout <-chan time.Time		// Создаю пиздец какой опасный nil chanal
	timeoutDur := time.NewTimer(dur * time.Second)		// ставлю таймер на dur в секундах
	timeout = timeoutDur.C		//передаю сигнал в канал 
	return timeout	  //Возвращаю канал
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
