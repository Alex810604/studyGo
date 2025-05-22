// 14. Найти индекс первого уникального символа
// Дана строка. Верните индекс первого символа, который встречается только один раз.
package main

import "fmt"

func main() {
	word := "ваппдднтоершва"

	if index := firstCharacter(word); index != -1 {
		fmt.Printf("Уникальный индекс первого символа '%c' равен %d.\n", []rune(word)[index], index)
	} else {
		fmt.Println("Уникальных индексов нет")
	}
}
func firstCharacter(word string) int {
	m := make(map[rune]int)
	word1 := []rune(word)
	for _, num := range word1 {
		m[num]++
	}
	for i, num := range word1 {
		if m[num] == 1 {
			return i
		}
	}
	return -1
}
