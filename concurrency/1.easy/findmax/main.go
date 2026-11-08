package main

import (
	"fmt"
	"sync"
)

// Что будет выведено? Почему? Подумать, после запустить программу. Исправить.

func main() {
	var maxi int
	var mu sync.Mutex
	var wg sync.WaitGroup

	for i := 1000; i > 0; i-- {
		wg.Add(1)
		go func(val int) {
			defer wg.Done()
			mu.Lock()
			if i%2 == 0 && val > maxi {
				maxi = val
			}
			mu.Unlock()
		}(i)
	}

	wg.Wait()
	fmt.Printf("Maximum is %d", maxi)
}

// Вывод: 1000. Нет сихронизации и мэин сразу закроется.
// Нужно добавить сихронизацию(для работы горутин) и мьютекс(Гонка данных)
