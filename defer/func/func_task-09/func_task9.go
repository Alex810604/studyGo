// 9. Группировка строк по длине
// Напишите функцию GroupByLength, которая принимает слайс строк и возвращает мапу,
// где ключ — длина строки, а значение — слайс строк этой длины.
package main

import (
	"fmt"
	"unicode/utf8"
)

func main() {
	word := []string{"книга", "альбом", "Иван", "класс", "привет", "цветок"}

	// вариант 1
	result1 := GroupByLength1(word)
	fmt.Println(result1)

	// вариант 2
	result2 := GroupByLength2(word)
	fmt.Println(result2)

	// вариант 3
	result3 := GroupByLength3(word, func(a string, m map[int][]string) {
		length := utf8.RuneCountInString(a)
		m[length] = append(m[length], a)
	})
	fmt.Println(result3)
}

// вариант 1
func GroupByLength1(str []string) map[int][]string {
	groups := make(map[int][]string)
	for _, segment := range str {
		length := utf8.RuneCountInString(segment)
		groups[length] = append(groups[length], segment)
	}
	return groups
}

func GroupByLength2(str []string) map[int][]string {
	groups := make(map[int][]string)

	addToGroup := func(segment string) {
		length := utf8.RuneCountInString(segment)
		groups[length] = append(groups[length], segment)
	}

	for _, segment := range str {
		addToGroup(segment)
	}
	return groups
}

// вариант 3
func GroupByLength3(str []string, d func(string, map[int][]string)) map[int][]string {
	groups := make(map[int][]string)
	for _, segment := range str {
		d(segment, groups)
	}
	return groups
}
