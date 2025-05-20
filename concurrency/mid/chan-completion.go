package main

import (
	"fmt"
	"math/rand"
	"reflect"
	"runtime"
	"strings"
	"time"
)

// Канал завершения
// Есть функция, которая произносит текст пословно (с некоторыми задержками):

func say(done chan<- struct{}, id int, text string) {
	for _, word := range strings.Fields(text) {
		fmt.Printf("Worker #%d says: %s...\n", id, word)
		dur := time.Duration(rand.Intn(100)) * time.Millisecond
		time.Sleep(dur)
	}
	done <- struct{}{}
	fmt.Printf("The Worker [%s] with id [%d] completed the work\n", GetFunctionName((say)), id)
}

func GetFunctionName(i any) string {
	return runtime.FuncForPC(reflect.ValueOf(i).Pointer()).Name()
}

// Запускаем несколько одновременных воркеров, по одной на каждую фразу:

func main() {
	phrases := []string{
		"go is awesome",
		"cats are cute",
		"rain is wet",
		"channels are hard",
		"floor is lava",
	}
	done := make(chan struct{}, len(phrases))
	for idx, phrase := range phrases {
		go say(done, idx+1, phrase)
	}
	for range phrases {
		<-done
	}
}

// Программа ничего не печатает — функция main() завершается до того, как отработает хотя бы один воркер:
// Использовать канал для завершения. Пролистай, если нужна подсказка.
//
//
//
//
//
//
// say(done chan<- struct{}, id int, phrase string).

