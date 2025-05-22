// 7. Объединение двух мап
// Даны две мапы:
// m1 := map[string]int{"a": 1, "b": 2}
// m2 := map[string]int{"b": 3, "c": 4}
// Создайте новую мапу, объединив их (при конфликте берётся значение из m2).
package main

import "fmt"

func main() {
	m1 := map[string]int{"a": 1, "b": 2}
	m2 := map[string]int{"b": 3, "c": 4}
	m3 := make(map[string]int)
	for k, v := range m1 {
		m3[k] = v
	}
	for k, v := range m2 {
		m3[k] = v
	}
	fmt.Println(m3)
}
