// 2. Порядок выполнения defer (LIFO)
// Напишите функцию с несколькими defer и покажите, что они выполняются в обратном порядке (LIFO — Last In, First Out).
package main

import "fmt"

func factorial(n int) int {
	if n == 0 {
		return 1
	}
	defer fmt.Print(n, " ")
	return n * factorial(n-1) // Рекурсивный случай: n! = n * (n-1)!
}

func main() {
	nums := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	for _, num := range nums {
		n1 := num
		defer fmt.Print(n1, " ")
	}
	fmt.Println(factorial(5)) // 120 (5! = 5*4*3*2*1 = 120)

}
