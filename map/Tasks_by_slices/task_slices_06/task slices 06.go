// 6. Объединение двух срезов
// Напишите функцию, которая принимает два среза и возвращает новый срез,
// содержащий все элементы обоих без использования append (на базе make и копирования).
// Пример:
// Вход: []int{1, 2}, []int{3, 4}
// Выход: [1 2 3 4]
package main

import "fmt"

func main() {
	slice1 := []int{1, 2}
	slice2 := []int{3, 4}
	sliceNow := make([]int, len(slice1)+len(slice2))
	copy(sliceNow, slice1)
	copy(sliceNow[len(slice1):], slice2)
	fmt.Println(sliceNow)

}
