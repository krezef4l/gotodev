package main

import (
	"errors"
	"time"
)

// Написать функцию after() — аналог time.After():

// возвращает канал, в котором появится значение      // какое значение будет передаваться я тоже не понялнихуя
// через промежуток времени dur
func after(dur time.Duration) <-chan time.Time { //Честно говоря я нихуя не понял, чё за тип переменной time.Duration 
//я несколько раз перечитал pkg.go.dev на эту тему, залез в исходный код языка и нихуя не понял.	
	var timeout = make(chan time.Time)  //Создаю канал, который буду возвращать
	defer close(timeout)		    //На всякий случай закрываю его по зевершению функции
	
	var myTime time.Tim		   //Создал переменную того типа который можно прокинуть по каналу
	time.Sleep(dur)			   //я думаю что сплю здесь но мне кажется это нихуя не так работает
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
