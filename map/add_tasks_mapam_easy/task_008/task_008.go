// 8. Подсчёт частоты символов в строке
// Дана строка "hello". Создайте мапу, где ключи — символы, а значения — сколько раз они встречаются.
package main

import "fmt"

func main() {
	str := "hello"
	m := make(map[string]int)
	for _, r := range str {
		m[string(r)]++
	}
	fmt.Println(m)

	for k, v := range m {
		fmt.Printf("Ключ-символ %s встречается %d раз\n", k, v)
	}
}
