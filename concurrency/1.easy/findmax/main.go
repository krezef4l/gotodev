package main

import (
	"fmt"
	"sync"
)

// Что будет выведено? Почему? Подумать, после запустить программу. Исправить.
/*изначально будет выведенно 1000 или другое значение, так как программма не правильно реализованна, функция main заканчивается раньше 
окончания работы горутины */

func main() {
	var maxi int
	var mu sync.Mutex
	var wg sync.WaitGroup

	for i := 1000; i > 0; i-- {
		wg.Add(1)

		go func(i int) {
			defer wg.Done()

			if i%2 == 0 {
				mu.Lock()

				if i > maxi {
					maxi = i
				}
				mu.Unlock()
			}

		}(i)
	}
	wg.Wait()

	fmt.Printf("Maximum is %d\n", maxi)
}
