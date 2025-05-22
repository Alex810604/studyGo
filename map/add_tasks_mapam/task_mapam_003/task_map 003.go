// 3. Поиск дубликатов в массиве
// Дан массив чисел. Верните true, если в массиве есть дубликаты, и false в противном случае.
package main

import "fmt"

func main() {
	array := [5]int{1, 5, 2, 3, 4}
	m := make(map[int]bool) // false  true
	k := false
	for _, v := range array {
		if m[v] {
			k = true
			break
		}
		m[v] = true
	}
	fmt.Println(k)
}
