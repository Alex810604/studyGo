// 10. Rate limiter на atomic
// Реализуйте простой ограничитель запросов (rate limiter), который:
// Считает количество запросов в текущем интервале
// Сбрасывает счетчик по таймеру
// Должен работать максимально эффективно при проверках (atomic).
package goroutine2

import (
	"fmt"
	"sync/atomic"
	"time"
)

type RateLimiter struct {
	limit    uint64        // лимит запросов на интервал
	count    atomic.Uint64 // текущий счётчик
	interval time.Duration
	stop     chan struct{}
}

func NewRateLimiter(limit uint64, interval time.Duration) *RateLimiter {
	return &RateLimiter{
		limit:    limit,
		interval: interval,
		stop:     make(chan struct{}),
	}
}
func (r *RateLimiter) Allow() bool {
	if r.count.Add(1) <= r.limit {
		return true
	}
	return false
}

func (r *RateLimiter) RunAutoReset() {
	go func() {
		ticker := time.NewTicker(r.interval)
		defer ticker.Stop()
		for {
			select {
			case <-ticker.C:
				r.count.Store(0) // Сброс счётчика запросов
			case <-r.stop:
				fmt.Println("Завершение сброса")
				return
			}
		}
	}()
}

// Завершение сброса
func (r *RateLimiter) Stop() {
	close(r.stop)
}
