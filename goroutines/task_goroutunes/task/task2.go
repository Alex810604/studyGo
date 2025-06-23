// 2. Множество горутин
// Запустите 5 горутин, каждая из которых выводит свой номер (от 1 до 5). Убедитесь, что все горутины успевают выполниться.
package goroutine

import (
	"fmt"
	"sync"
)

func LotsGoroutines(id int, wg *sync.WaitGroup) {
	defer wg.Done()
	fmt.Printf("Горутина № %d\n", id)
}
