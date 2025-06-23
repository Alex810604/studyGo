//4. Горутины и цикл
//Запустите 3 горутины внутри цикла, передавая каждой значение счётчика (i). Убедитесь, что вывод соответствует ожиданиям (0, 1, 2).

package goroutine

import (
	"fmt"
	"sync"
)

func CycleGoroutine(id int, wg *sync.WaitGroup) {
	defer wg.Done()
	fmt.Println("значение счётчика ", id)
}
