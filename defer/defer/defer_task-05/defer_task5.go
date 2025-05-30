// 5. Defer внутри цикла.
// Напишите функцию с циклом, где defer вызывается на каждой итерации. Объясните, почему все defer выполняются после цикла.
package main

import (
	"fmt"
)

func main() {
	for i := 0; i < 5; i++ {
		defer fmt.Print(i, " ")
	}
}
