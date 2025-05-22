// 15. Подсчет количества подстрок с одинаковыми символами на концах
// Дана строка. Посчитайте количество подстрок, которые начинаются и заканчиваются одним и тем же символом.
package main

import "fmt"

func main() {
	word := "пдд"
	m := make(map[string]int)
	for _, n := range word {
		m[string(n)]++
	}
	fmt.Println(m)
	total := 0
	for _, v := range m {
		total += v * (v + 1) / 2
	}
	fmt.Println(total)
}
