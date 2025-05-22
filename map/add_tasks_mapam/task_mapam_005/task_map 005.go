// 5. Проверка анаграмм
// Даны две строки. Проверьте, являются ли они анаграммами (состоят из одних и тех же символов в разном порядке).
package main

import "fmt"

func main() {
	str1 := "Павел левак в поле ковал"
	str2 := "в поле Павел ковал левак"
	n := ""

	m1 := anagrams(str1)
	m2 := anagrams(str2)
	if len(m1) != len(m2) {
		n = "Cтроки не являются анаграммами"

	} else {
		for k, v := range m1 {
			if m2[k] != v {
				n = "Cтроки не являются анаграммами"
				break
			}
		}
	}
	if n == "" {
		n = "Cтроки являются анаграммами"
	}
	fmt.Println(n)
}
func anagrams(str string) map[rune]int {
	m := make(map[rune]int)
	for _, v := range str {
		m[v]++
	}
	return m
}
