package main

import (
	"fmt"
	"math/rand"
	"strings"
	"time"
)

// say произносит текст пословно с некоторыми задержками.
func say(done chan<-struct{}, id int, text string) {
	defer func(){
		done<-struct{}{}
	}()
	for _, word := range strings.Fields(text) {
		fmt.Printf("Worker #%d says: %s...\n", id, word)
		dur := time.Duration(rand.Intn(100)) * time.Millisecond
		time.Sleep(dur)
	}
}

func main() {
	done := make(chan struct{})
	phrases := []string{
		"go is awesome",
		"cats are cute",
		"rain is wet",
		"channels are hard",
		"floor is lava",
	}

	// Запускаем несколько одновременных воркеров, по одной на каждую фразу.
	for idx, phrase := range phrases {
    go say(done, idx+1, phrase)
}
		
	for range phrases {
    <-done

	}
}

// Эта программа ничего не напечатает — функция main() завершается до того, как
// отработает хотя бы один воркер. Чтобы исправить, нужно использовать канал для
// завершения. Пролистай ниже, если понадобится подсказка.
//
//
//
//
//
//
//
//
//
//
//
// Так как знаем количество строк заранее, можно использовать 1 канал
// завершения, из которого будем вычитывать значение N раз (где N - количество
// строк). Канал нужно послать как аргумент функции say:
//
// say(done chan<- struct{}, id int, phrase string).
//
// + Подумать как быть, если бы количество строк не было известно заранее.
