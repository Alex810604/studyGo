// 8. Найти самый частый элемент
// Дан массив чисел. Найдите число, которое встречается чаще всего.
package main

import "fmt"

func main() {
	n1 := []int{4, 5, 8, 4, 8, 4}
	m := make(map[int]int)
	for _, v := range n1 {
		m[v]++
	}
	vMax := 0
	iMax := 0
	for i, v := range m {
		if v > vMax {
			vMax = v
			iMax = i
		}
	}
	fmt.Print(iMax)
}
