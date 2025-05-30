// 2. Порядок выполнения defer (LIFO)
// Напишите функцию с несколькими defer и покажите, что они выполняются в обратном порядке (LIFO — Last In, First Out).
package main

import (
	"fmt"
)

func main() {
	nums := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	for _, num := range nums {
		defer fmt.Print(num, " ") //10 9 8 7 6 5 4 3 2 1
	}
	defer fmt.Println()
	for i := 0; i < 5; i++ {
		defer fmt.Print(i, " ") //4 3 2 1 0
	}
	defer fmt.Println()
	for j := 6; j < 10; j++ {
		defer func(n int) {
			fmt.Print(n, " ") //9 8 7 6
		}(j)
	}
}
