// 8. Graceful shutdown флаг
// Реализуйте механизм graceful shutdown, где:
// Одна горутина может установить флаг завершения
// Множество горутин могут проверять этот флаг
// Требуется максимальная производительность при проверках (atomic).
package goroutine2

import "sync/atomic"

// GracefulShutdown предоставляет атомарный флаг
type GracefulShutdown struct {
	flag atomic.Bool
}

// SetShutdown устанавливает флаг завершения true (запись)
func (g *GracefulShutdown) SetShutdown() {
	g.flag.Store(true)
}

// Проверка флага (чтение)
func (g *GracefulShutdown) GetShutdown() bool {
	return g.flag.Load()
}
