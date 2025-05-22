// 1. Подсчет частоты символов
// Дана строка. Посчитайте, сколько раз каждый символ встречается в строке.
package main

import (
	"fmt"
)

func main() {
	str := "Павел Павлин Петр"
	m := make(map[rune]int)
	for _, n := range str {
		m[n] += 1
	}
	m2 := make(map[string]int)
	for i, n := range m {
		m2[string(i)] = n

	}
	fmt.Println(m)
	fmt.Println(m2)
}
