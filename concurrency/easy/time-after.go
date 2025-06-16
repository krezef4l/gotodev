package main

import (
	"errors"
	"time"
)

// Написать функцию after() — аналог time.After():

// возвращает канал, в котором появится значение      // какое значение будет передаваться я тоже не понял нихуя
// через промежуток времени dur
func after(dur time.Duration) <-chan time.Time {
	var timeout = make(chan time.Time)  //Создаю канал, который буду возвращать
	defer close(timeout)		    //На всякий случай закрываю его по зевершению функции
	
	var myTime time.Time		   //Создал переменную того типа который можно прокинуть по каналу
	time.Sleep(dur * time.Second       //я думаю что сплю 10 сек здесь? но мне кажется это нихуя не так работает
	timeout<- myTime		   //Прокидываю по каналу переменную которую можно прокинуть
	return timeout			  //Возвращаю канал
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
