package main

import "fmt"

// Не запуская, подумать, что будет выведено на экран.
// Проверить и разобраться, почему так.
// Написать ответ подробно.

func main() {
	x := []int{}
	x = append(x, 0) // count=0,len(1),cap(1)
	fmt.Println("count: x", x, "len", len(x), "cap", cap(x))

	x = append(x, 1) // count = 0, 1, len(2),cap(2)
	fmt.Println("count: x", x, "len", len(x), "cap", cap(x))

	x = append(x, 2) // count = 0, 1, 2, len(3), cap(4)
	fmt.Println("count: x", x, "len", len(x), "cap", cap(x))

	y := append(x, 3) // count = 0,1,2,3, len(4),cap,(4)
	fmt.Println("count: y", y, "len", len(y), "cap", cap(y))

	z := append(x, 4) // count = 0,1,2,4, len(4),cap(4)
	fmt.Println("count: z", z, "len", len(z), "cap", cap(z))
	fmt.Println("res:", y, z)
	fmt.Println("адрес:", &x[0], &y[0], &z[0])

	// чтобы получить разные значение, создает разные слайсы
	u := make([]int, len(x))
	copy(u, x)
	u = append(u, 3)

	s := make([]int, len(x))
	copy(s, x)
	s = append(s, 4)

	fmt.Println("u:", u)
	fmt.Println("s:", s)
}

/*
y = 0 1 2 4, z = 0 1 2 4.
оба заголовка укзаывают на оджин массив.
можно через адрес.
Но
*/
