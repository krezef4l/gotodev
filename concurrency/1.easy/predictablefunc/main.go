package main

import (
	"fmt"
	"context"
	"errors"
	"math/rand"
	"time"
)

func init() {
	rand.Seed(time.Now().UnixNano())
}

// Есть функция, работающая неопределённо долго и возвращающая число. Её тело
// нельзя изменять (представим, что внутри сетевой запрос).
func unpredictableFunc() int64 {
	rnd := rand.Int63n(5000)
	time.Sleep(time.Duration(rnd) * time.Millisecond)
	return rnd
}


func unpredictableFuncWithTimeout()(int64,error){
	start:=time.Now()
	defer func(){
		fmt.Println("Execution time:",time.Since(start))
	}()
	
ctx, cancel:=context.WithTimeout(context.Background(), 2*time.Second)
	defer cancel()
	resultCh:=make(chan int64,1)
	go func (){
		result:=unpredictableFunc()
		resultCh<-result
	}()
	select{
case result:=<-resultCh:
		return result, nil
		case <-ctx.Done():
return 0,errors.New("Timeout")
		}
	
}
// Нужно изменить функцию-обёртку, которая будет работать с заданным таймаутом
// (например, 1 секунду). Если "длинная" функция отработала за это время -
// отлично, возвращаем результат. Если нет - возвращаем ошибку. Результат работы
// в этом случае нам не важен.
//
// Дополнительно нужно измерить, сколько выполнялась эта функция (просто вывести
// в лог). Сигнатуру функцию обёртки менять можно.

func predictableFunc() int64 {
	return unpredictableFunc()
}

func main() {
	res, err := unpredictableFuncWithTimeout()
	if err !=nil{
		fmt.Println("Error:",err)
		return
	}
	fmt.Println("Result",res)
}
