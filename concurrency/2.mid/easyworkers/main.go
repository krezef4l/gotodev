package main

import (
  "fmt"
  "sync"
)

func worker(id int, jobs <-chan int, results chan<- int) {
  for num := range jobs {
  fmt.Println("Worker %d received: %d\n", id, num)
      time.Sleep(500 * time.Millisecond)
      result := num * num
      results <- result
  }
}

func main() {
var wg sync.WaitGroup
    jobs := make(chan int)
    results := make(chan int, 10)

    // Запускаем 3 воркера
    for i := 1; i <= 3; i++ {
    wg.Add(1)

    go func(id int) {
        defer wg.Done()
        worker(id, jobs, results)
    }(i)
}

    // Отправляем 10 задач
    for i := 1; i <= 10; i++ {
        jobs <- i
    }
    close(jobs)
go func() {
    wg.Wait()
    close(results)
}()
    // Получаем результаты
  for result := range results {
    fmt.Println("Result:", result)
}
}
```
// Реализовать пул из 3-х воркеров, которые:
// - получают задачи (в задачах просто спим и что-то печатаем, например) из
// общего канала;
// - вычисляют квадрат числа и отправляют результат в общий канал.
// Главная горутина создаёт N задач, распределяет их по воркерам и выводит
// результаты.
