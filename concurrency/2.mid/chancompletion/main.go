package main

import (
	"fmt"
	"math/rand"
	"strings"
	"time"
)

// say произносит текст пословно с некоторыми задержками.
func say(id int, text string) {
	for _, word := range strings.Fields(text) {
		fmt.Printf("Worker #%d says: %s...\n", id, word)
		dur := time.Duration(rand.Intn(100)) * time.Millisecond
		time.Sleep(dur)
	}
}

func main() {
	phrases := []string{
		"go is awesome",
		"cats are cute",
		"rain is wet",
		"channels are hard",
		"floor is lava",
	}
	for idx, phrase := range phrases {
		// Запускаем несколько одновременных воркеров, по одной на каждую фразу.
		go say(idx+1, phrase)
	}
}

// Программа ничего не напечатает — функция main() завершается до того, как отработает хотя бы один воркер.
// Чтобы исправить, нужно использовать канал для завершения. Пролистай ниже, если понадобится подсказка.
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
// Так как знаем количество строк заранее, можно использовать 1 канал завершения, из которого будем вычитывать
// значение N раз (где N - количество строк). Канал нужно послать как аргумент функции say:
// say(done chan<- struct{}, id int, phrase string).
//
// + Подумать как быть, если бы количество строк не было известно заранее.
