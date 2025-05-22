// 2. Подсчет частоты слов
// Дана строка с словами. Посчитайте, сколько раз каждое слово встречается в строке.
package main

import "fmt"

func main() {
	str := []string{"Павел", "Павлин", "Павлин", "Петр", "Петр", "Петр"}
	m := make(map[string]int)
	for _, v := range str {
		m[v]++
	}
	fmt.Println(m)
}
