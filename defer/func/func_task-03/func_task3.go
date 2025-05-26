// 3. Объединение двух мап
// Напишите функцию MergeMaps, которая принимает две мапы (map[string]int) и возвращает новую мапу, объединяющую их.
// Если ключ присутствует в обеих мапах, выбирается значение из второй.
package main

import "fmt"

func main() {
	m1 := map[string]int{"отдых": 1, "респект": 3, "дом": 1}
	m2 := map[string]int{"респект": 2, "книга": 2, "Ok": 2}
	fmt.Println(MergeMaps(m1, m2))
}
func MergeMaps(d1, d2 map[string]int) map[string]int {
	d3 := make(map[string]int)
	count := func(d map[string]int) {
		for i, v := range d {
			d3[i] = v
		}
	}
	count(d1)
	count(d2)
	return d3
}
