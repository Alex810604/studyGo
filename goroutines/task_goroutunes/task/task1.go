//1. Простая горутина
//Напишите программу, которая запускает горутину, выводящую "Hello, Go!", и дожидается её завершения с помощью time.Sleep.

package goroutine

import (
	"fmt"
)

func SimpleGoroutine() {
	fmt.Println("Hello, Go!")
}
