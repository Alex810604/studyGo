// 10. Проверить, являются ли строки изоморфными
// Две строки называются изоморфными, если можно заменить символы одной строки на символы другой (с сохранением порядка).
// Например, "egg" и "add" изоморфны.
package main

import "fmt"

func main() {
	word1, word2 := "Add", "пдд"
	fmt.Println(isomorphic(word1, word2))
	//fmt.Println([]rune(word1), []rune(word2))
}
func isomorphic(word1, word2 string) bool {
	runes1 := []rune(word1)
	runes2 := []rune(word2)
	if len(runes1) != len(runes2) {
		return false
	}
	m1 := make(map[rune]rune)
	m2 := make(map[rune]rune)
	for i := 0; i < len(runes1); i++ {
		number1 := runes1[i]
		number2 := runes2[i]

		if value, exists := m1[number1]; exists {
			if value != number2 {
				return false
			}
		} else {
			m1[number1] = number2
		}
		if value, exists := m2[number2]; exists {
			if value != number1 {
				return false
			}
		} else {
			m2[number2] = number1
		}
	}
	return true
}
