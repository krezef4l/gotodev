package main

import (
	"fmt"
)

// Написать функцию double, которая работает со слайсом указателей на int:
// - функция должна умножать все уникальные значения на 2 (slice 3 3 3-> 6 6 6);
// - естественно, в слайсе будут повторяющиеся значения (адреса). Повторно
// умножать нельзя;
// - inplace реализация: не перекладываем из слайса в слайс.

func double(numbers []*int) {
	//TODO:
	mmap := make(map[*int]struct{}, len(numbers))

	for _, p := range numbers {
		if p == nil {
			continue
		}

		if _, ok := mmap[p]; ok {
			continue
		}
		mmap[p] = struct{}{}
		*p *= 2
	}
}

func main() {
	numbers := make([]*int, 0, 5)
	var number int
	for range 3 {
		number++
		numbers = append(numbers, &number)
	}

	for _, number := range numbers {
		fmt.Printf("%d ", *number)
	}
	fmt.Println()

	// double(numbers)

	// for _, number := range numbers {
	// 	fmt.Printf("%d ", *number)
	// }
	// fmt.Println()
}
