// 16. Проверить, можно ли получить строку перестановкой другой
// Даны две строки. Проверьте, можно ли получить одну из другой перестановкой символов.
package main

import "fmt"

func main() {
	word1 := "дпд"
	word2 := "ппп"
	result := symbolPermutation(word1, word2)
	fmt.Printf("'%s' и '%s' -> %v.\n", word1, word2, result)
}
func symbolPermutation(word1, word2 string) bool {
	m := make(map[string]int)
	if len(word1) != len(word2) {
		return false
	}
	for _, n := range word1 {
		m[string(n)]++
	}
	for _, n := range word2 {
		m[string(n)]--
		if m[string(n)] < 0 {
			return false
		}
	}
	return true
}
