// 2. Объединение двух мап
// Напишите функцию, которая принимает две map[string]int и возвращает новую мапу, объединяя их.
// Если ключ присутствует в обеих мапах, сложите значения.
// Пример:
// Вход:
// m1 := map[string]int{"a": 1, "b": 2}
// m2 := map[string]int{"b": 3, "c": 4}
// Выход: map[a:1 b:5 c:4]
package main

import "fmt"

func main() {
	m1 := map[string]int{"a": 1, "b": 2}
	m2 := map[string]int{"b": 3, "c": 4}
	m3 := mapMerge(m1, m2)
	fmt.Println(m3)
}
func mapMerge(m1, m2 map[string]int) map[string]int {
	m3 := make(map[string]int)
	for i, _ := range m1 {
		m3[i] = m1[i]
	}
	fmt.Println(m3)
	for i, _ := range m2 {
		m3[i] += m2[i]
	}
	return m3
}
