package main

import (
	"context"
	"fmt"
)

// Что будет, если один и тот же контекст отменить дважды. Почему?
/*ничего не произойдет кроме текстового вывода, так  как после первого кенсел, второй не изменит ничего в программе*/
// Подумать и только после этого запустить программу.

func main() {
	ctx, cancel := context.WithCancel(context.Background())

	cancel()
	fmt.Println("Отменено 1 раз:", ctx.Err())

	cancel()
	fmt.Println("???", ctx.Err())
}
