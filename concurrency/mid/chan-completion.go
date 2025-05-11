package main

import (
	"fmt"
	"math/rand"
	"strings"
	"time"
)

// Канал завершения
// Есть функция, которая произносит текст пословно (с некоторыми задержками):

func say(done chan struct{}, id int, text string) {
	for _, word := range strings.Fields(text) {
		fmt.Printf("Worker #%d says: %s...\n", id, word)
		dur := time.Duration(rand.Intn(100)) * time.Millisecond
		time.Sleep(dur)
	}

	close(done)
}

// Запускаем несколько одновременных воркеров, по одной на каждую фразу:

func do6() {
	phrases := []string{
		"go is awesome",
		"cats are cute",
		"rain is wet",
		"channels are hard",
		"floor is lava",
	}

	channels := make([]chan struct{}, len(phrases))

	for i := range channels {
		channels[i] = make(chan struct{})
	}

	for idx, phrase := range phrases {
		go say(channels[idx], idx+1, phrase)
	}

	for _, ch := range channels {
		<-ch
	}
}

// Программа ничего не печатает — функция main() завершается до того, как отработает хотя бы один воркер:
// Использовать канал для завершения. Пролистай, если нужна подсказка.