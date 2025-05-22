// 6. Преобразование мапы в срез ключей/значений
// Напишите функцию, которая принимает map[int]string и возвращает два среза: один с ключами, другой со значениями.
// Пример:
// Вход: map[int]string{1: "a", 2: "b"}
// Выход: []int{1, 2}, []string{"a", "b"}
package main

import "fmt"

func main() {
	db := map[int]string{1: "a", 2: "b"}
	key, meanings := sliceInteger(db)
	fmt.Print(key, meanings)
}
func sliceInteger(db map[int]string) ([]int, []string) {
	key := make([]int, 0, len(db))
	meanings := make([]string, 0, len(db))
	for i, n := range db {
		key = append(key, i)
		meanings = append(meanings, n)
	}
	return key, meanings
}
