package main

import "fmt"

// Что будет выведено? Почему? Подумать, после запустить программу.

func main() {
	sl := make([]string, 5, 7)
	f1(sl)
	fmt.Println(sl)
	fmt.Println(sl[:6]) // для проверки
}

func f1(sl []string) {
	sl = append(sl, "Hello World!")
}

/* вывод будет 5 пустых индексов [    ]
но почему не виден 6 индекс? из за старого заголовка он не видит!
а фактический он там есть

*/
