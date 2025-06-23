// 6. Исправление гонки через sync.Mutex
// Модифицируйте предыдущую задачу, используя sync.Mutex, чтобы избежать гонки данных.
package goroutine

import (
	"fmt"
	"sync"
)

func DataRaceMutex(wg *sync.WaitGroup, mu *sync.Mutex, id int, counter *int) {
	defer wg.Done()
	mu.Lock()
	*counter++
	fmt.Println("Горутина", id, "увеличила счетчик на ", *counter)
	mu.Unlock()
}
