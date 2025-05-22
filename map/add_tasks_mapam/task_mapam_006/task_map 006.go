// 6. Найти пересечение двух массивов
// Даны два массива чисел. Верните массив их пересечения (общие элементы).
package main

import "fmt"

func main() {
	n1 := [6]int{1, 3, 3, 3, 4, 5}
	n2 := [5]int{8, 3, 2, 3, 5}
	m := make(map[int]int)

	for _, v := range n1 {
		m[v]++
	}
	fmt.Println(m)
	var result []int
	for _, num := range n2 {
		if m[num] > 0 {
			fmt.Println(m)
			result = append(result, num)
			m[num]--
		}
	}
	fmt.Println(m, result)
}
