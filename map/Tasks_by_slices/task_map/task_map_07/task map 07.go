// 7. Удаление дубликатов из среза с помощью мапы
// Напишите функцию, которая принимает срез строк и возвращает новый срез без дубликатов,
// используя мапу для проверки уникальности.
// Пример:
// Вход: []string{"a", "b", "a", "c"}
// Выход: []string{"a", "b", "c"}
package main

import "fmt"

func main() {
	str := []string{"a", "b", "a", "c"}
	str2 := duplicate(str)
	fmt.Print(str2)
}
func duplicate(str []string) []string {
	db := make(map[string]struct{})
	var result []string
	for _, n := range str {
		if _, exists := db[n]; !exists {
			db[n] = struct{}{}
			result = append(result, n)
		}
	}
	return result
}
