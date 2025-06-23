// 6. Атомарный генератор ID
// Реализуйте генератор уникальных идентификаторов, который гарантированно возвращает
// // уникальные числа даже при вызове из множества горутин. Используйте atomic.
package goroutine2

import "sync/atomic"

type AtomicGenerator struct {
	count int64
}

func NewAtomicGenerator() *AtomicGenerator {
	return &AtomicGenerator{count: 0}
}
func (a *AtomicGenerator) ReturnUniqueNumbers() int64 {
	return atomic.AddInt64(&a.count, 1)
}
