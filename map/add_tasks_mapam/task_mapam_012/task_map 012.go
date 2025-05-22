// 12. Найти пропущенное число
// Дан массив чисел от 1 до n с одним пропущенным числом. Найдите его, используя хешмапу.
package main

import (
	"fmt"
)

func main() {
	nums := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 12, 13, 14, 15, 16, 17, 18, 19, 20}
	n := 20
	m := make(map[int]struct{})
	total := 0
	for _, b := range nums {
		m[b] = struct{}{}
	}
	for i := 1; i <= n; i++ {
		if _, ok := m[i]; !ok {
			total = i
			break
		}
	}
	fmt.Println(m)
	fmt.Printf("Пропущенное число: %d.\n", total)
}
