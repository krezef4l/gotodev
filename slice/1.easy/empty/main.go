package main

import "fmt"

// Не запуская, подумать, что будет выведено на экран (true/false).
// Проверить и разобраться, почему так.
// Написать ответ подробно.

func main() {
	var list []int
	fmt.Println(list == nil) //true

	list = []int{}
	fmt.Println(list == nil) //fasle

}

/*Как я понял, тут можно сказать 2 true,
Но результат true false, Почему первый вариант реально nil
где он равен len=0,cap=0 list==nil=true,
а []list{} len=0,cap=0, list==nil=false,
[]list[]-non-nil с пустными набором элементов,
с указанием на служебную область runtime.zerobase
*/
