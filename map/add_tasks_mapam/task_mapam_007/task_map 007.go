// 7. Подсчет суммы двух чисел
// Дан массив чисел и целевое число. Проверьте, есть ли в массиве два числа, сумма которых равна целевому.
package main

import "fmt"

func main() {
	n1 := []int{1, 2, 3, 4, 5, 6}
	n := 11
	fmt.Print(sumNumbers(n1, n))
}
func sumNumbers(n1 []int, n int) bool {
	m := make(map[int]bool)
	//flag := false
	for _, v := range n1 {
		vNow := n - v
		if m[vNow] {
			return true
		}
		m[v] = true
	}
	return false
}
