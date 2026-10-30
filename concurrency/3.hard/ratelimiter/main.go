package main

import (
	"context"
	"errors"
	"fmt"
	"sync"
	"time"
)

// Есть функция, которая выполняет http-запрос. Необходимо реализовать логику
// rate limiter с возможностью ограничения кол-ва запросов, производимых этой
// функцией до n-го кол-ва в секунду.

// общий лимитер и общий контекст
var (
	rl  *RateLimiter
	ctx context.Context
)

// структура - обёртка для time.Ticker
type RateLimiter struct {
	ticker *time.Ticker
}

// Конструктор n-запросов / сек
// Используем Ticker.C для отправки значений через канал(потокобезопасен)
func NewRateLimiter(n int) *RateLimiter {
	return &RateLimiter{
		ticker: time.NewTicker(time.Second / time.Duration(n)),
	}
}

// Wait - Блокируется до одного из двух событий
func (rl *RateLimiter) Wait(ctx context.Context) error {
	select {
	case <-rl.ticker.C:
		return nil
	case <-ctx.Done():
		return ctx.Err()
	}
}

// Close - освобождения ресурса
func (rl *RateLimiter) Close() {
	rl.ticker.Stop()
}

func makeRequest(url, data string) error {
	// Здесь могла быть реальная работа HTTP-клиента.
	fmt.Println("Request sent to:", url, "with data:", data, "at",
		time.Now().Format("15:04:05.000"))
	return nil
}

// обёртка с лимитом
func makeRequestWithLimit() error {
	if err := rl.Wait(ctx); err != nil {
		return err
	}
	return makeRequest("https://api.example.com/hotel", "query=123")

}

func main() {
	rl = NewRateLimiter(3)
	defer rl.Close()

	var cancel context.CancelFunc
	ctx, cancel = context.WithTimeout(context.Background(), 2*time.Second)
	defer cancel()

	var wg sync.WaitGroup
	for i := 0; i < 10; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			if err := makeRequestWithLimit(); err != nil {
				if errors.Is(err, context.DeadlineExceeded) {
					fmt.Println("skip: deadline execeeded")
				} else {
					fmt.Println("error:", err)
				}
			}
		}()
	}
	wg.Wait()
}
