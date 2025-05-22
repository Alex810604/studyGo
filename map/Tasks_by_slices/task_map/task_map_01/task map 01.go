// 1. Подсчет частоты элементов.
// Напишите функцию, которая принимает срез строк и возвращает map[string]int,
// где ключ — это строка, а значение — количество её вхождений.
// Пример:
// Вход: []string{"a", "b", "a", "c", "b"}
// Выход: map[a:2 b:2 c:1]
package main

import "fmt"

func main() {
	str := []string{"a", "b", "a", "c", "b"}
	result := CountElements(str)
	fmt.Println(result)
}
func CountElements(str []string) map[string]int {
	strNow := make(map[string]int)
	for _, v := range str {
		strNow[v]++
	}
	return strNow
}
