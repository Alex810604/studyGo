// 8. Разворот среза.
// Напишите функцию, которая разворачивает срез in-place (первый элемент становится последним и т.д.).
// Пример:
// Вход: []int{1, 2, 3, 4}
// Выход: [4 3 2 1]
package main

import "fmt"

func main() {
	slice := []int{1, 2, 3, 4, 5}
	sliceSpread := make([]int, 0, len(slice))
	for i := len(slice) - 1; i >= 0; i-- {
		sliceSpread = append(sliceSpread, slice[i])
	}
	fmt.Println(sliceSpread)
	//или без создания нового среза
	slice = reverseSlice(slice)
	fmt.Println(slice)
}

func reverseSlice(slice []int) []int {
	for i, j := 0, len(slice)-1; i < j; i, j = i+1, j-1 {
		slice[i], slice[j] = slice[j], slice[i] // Меняем местами
	}
	return slice
}
