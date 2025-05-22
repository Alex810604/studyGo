// 11. Подсчет гласных в строке
// Дана строка. Посчитайте количество каждой гласной буквы (a, e, i, o, u).
package main

import (
	"fmt"
	"strings"
)

func main() {
	word := "AJkOIOEx,apprejiuoiua"
	word = strings.ToLower(word)
	vowels := "aeiouy"
	m := make(map[string]int)
	for _, v := range word {
		v1 := string(v)
		for _, vowel := range vowels {
			if v1 == string(vowel) {
				m[v1]++
				break
			}
		}
	}
	fmt.Println(m)
}
