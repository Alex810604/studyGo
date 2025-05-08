// Напишите функцию Increment(n *int), которая увеличивает значение переданного числа на 1.
package main

import "fmt"

func main() {
	var num int
	fmt.Printf("Введите значение num: ")
	fmt.Scan(&num)
	Increment(&num)
	fmt.Printf("Увеличенное значение 'num' на 1 равно: %d\n", num)
}
func Increment(n *int) {
	*n += 1
}
