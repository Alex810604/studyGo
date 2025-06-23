// 10. Завершение горутин по флагу.
//Создайте горутину, которая работает в бесконечном цикле, пока не будет установлен флаг (через атомарную переменную atomic.Bool)

package goroutine

import (
	"fmt"
	"sync"
	"sync/atomic"
	"time"
)

func InfiniteLoop(wg *sync.WaitGroup, flag *atomic.Bool) {
	defer wg.Done()
	for {
		if flag.Load() {
			fmt.Println("Горутина завершает работу")
			return
		}
		fmt.Println("Работает...")
		time.Sleep(500 * time.Millisecond)

	}

}
