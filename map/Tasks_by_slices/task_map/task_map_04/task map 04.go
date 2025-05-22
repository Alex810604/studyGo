// 4. Проверка на анаграммы
// Напишите функцию, которая проверяет, являются ли две строки анаграммами
// (содержат одинаковые символы в разном порядке), используя мапу для подсчета символов.
// Пример:
// Вход: "listen", "silent"
// Выход: true
package main

import "fmt"

func main() {
	str1, str2 := "listen", "silent"
	m1 := strMap(str1)
	fmt.Println(m1)
	m2 := strMap(str2)
	fmt.Println(m2)
	result := mapAnagrame(m1, m2)
	fmt.Println(result)
}
func strMap(str string) map[rune]int {
	m := make(map[rune]int)
	for _, n := range str {
		m[n]++
	}
	return m
}

func mapAnagrame(m1, m2 map[rune]int) bool {
	if len(m1) != len(m2) {
		return false
	}
	for i, n := range m1 {
		if m2[i] != n {
			return false
		}
	}
	return true
}
