//3. Горутина с задержкой.
//Создайте горутину, которая ждёт 2 секунды (time.Sleep), а затем выводит "Done". Главный поток должен дождаться её завершения.

package goroutine

import (
	"fmt"
	"sync"
	"time"
)

func GoroutineDelay(wg *sync.WaitGroup) {
	defer wg.Done()
	time.Sleep(time.Second * 2)
	fmt.Println("Done")

}
