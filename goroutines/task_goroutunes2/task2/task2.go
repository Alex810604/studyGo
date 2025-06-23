// 2. Атомарный счетчик
// Создайте потокобезопасный счетчик, который может инкрементироваться и
// возвращать текущее значение, используя только атомарные операции (atomic).
package goroutine2

import (
	"sync/atomic"
)

type CounterAtomic struct {
	value int64
}

func (c *CounterAtomic) Increment() {
	atomic.AddInt64(&c.value, 1)
}

func (c CounterAtomic) CurrentValue() int64 {
	return atomic.LoadInt64(&c.value)
}
