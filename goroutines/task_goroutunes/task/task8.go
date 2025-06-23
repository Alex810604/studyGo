// 8. Горутины и возврат значений
// Создайте функцию, которая запускает горутину, вычисляющую квадрат числа,
// и возвращает результат через разделяемую переменную (с синхронизацией).

//		AddInt32(&var, delta) — атомарно увеличивает переменную var на delta.
//	🔹 StoreInt32(&var, value) — атомарно записывает value в var.
//	🔹 LoadInt32(&var) — атомарно читает значение var.
//	🔹 CompareAndSwapInt32(&var, old, new) — атомарно изменяет var, если её текущее значение равно old.

package goroutine

import (
	"sync"
	"sync/atomic"
)

func NumberSquared(wg *sync.WaitGroup, result *int64, id int) {
	defer wg.Done()
	squared := int64(id * id)
	atomic.StoreInt64(result, squared)
}
