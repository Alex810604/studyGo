// 7. Ожидание горутин с sync.WaitGroup
// Запустите 4 горутины, каждая из которых увеличивает счётчик, и дождитесь их завершения с помощью sync.WaitGroup.
package goroutine

import (
	"fmt"
	"sync"
)

func DataWait(wg *sync.WaitGroup, id int, counter *int) {
	defer wg.Done()
	*counter++
	fmt.Println("Горутина", id, "увеличила счетчик на ", *counter)
}
