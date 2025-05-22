// 3. Фильтрация мапы по значению.
// Напишите функцию, которая принимает map[string]int и возвращает новую мапу,
// содержащую только элементы со значениями больше заданного числа n.
// Пример:
// Вход: map[string]int{"a": 10, "b": 5, "c": 20}, n = 15
// Выход: map[c:20]
package main

import "fmt"

func main() {
	m1 := map[string]int{"a": 10, "b": 5, "c": 20}
	n := 15
	m2 := mapMerge(n, m1)
	fmt.Println(m2)
}
func mapMerge(n int, m1 map[string]int) map[string]int {
	m2 := make(map[string]int)
	for i, _ := range m1 {
		if m1[i] > n {
			m2[i] = m1[i]
		}
	}
	return m2
}
