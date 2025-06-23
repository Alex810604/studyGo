// 5. Гонка данных
// Создайте программу с общей переменной, которую увеличивают несколько горутин.
// Покажите, что без синхронизации возможна гонка данных.
package goroutine

import (
	"fmt"
	"sync"
)

func DataRace(wg *sync.WaitGroup, id int, counter *int) {
	defer wg.Done()
	*counter++
	fmt.Println("Горутина ", id, "увеличила счетчик на ", *counter)
}
