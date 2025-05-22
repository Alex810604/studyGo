// 18. Подсчет количества пар с заданной суммой
// Дан массив чисел и число k. Посчитайте количество пар чисел, сумма которых равна k.
package main

import "fmt"

func main() {
	num := []int{1, 3, 5, 7, 8, 9, 11}
	k := 8
	d1 := pairsNumbers(num, k)
	fmt.Println(d1)
}
func pairsNumbers(num []int, k int) int {
	m1 := make(map[int]int)
	total := 0
	for _, n := range num {
		index := k - n
		if _, ok := m1[index]; ok { // или запись общей строкой: total += m1[index]
			total++ //
		}
		m1[n]++
	}

	return total
}
