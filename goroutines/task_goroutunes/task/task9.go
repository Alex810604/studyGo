// 9. Паника в горутине
//Напишите программу, где одна из горутин вызывает panic, и обработайте её с помощью recover в главном потоке.

package goroutine

import (
	"fmt"
	"sync"
	"sync/atomic"
)

func NumberCubed(wg *sync.WaitGroup, n int, result *int32) {
	defer wg.Done()
	cubed := int32(n * n * n)
	atomic.StoreInt32(result, cubed)
	panic("cбой работы")
	fmt.Println("Все что после 'panic' не выполнится")

}
