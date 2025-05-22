// 2. Изменение in-place
// Напишите функцию, которая принимает срез целых чисел и умножает каждый его элемент на 2 без возврата нового среза
// (изменя исходный).
// Пример:
// Вход: []int{1, 2, 3}
// Выход: [2 4 6]
package main

import "fmt"

func main() {
	fmt.Print("Введите целое положительное число n: ")
	var n int
	fmt.Scan(&n)
	original := make([]int, n)
	for i := 0; i < n; i++ {
		original[i] = i + 1
	}
	fmt.Printf("Вход: %d.\n", original)
	fmt.Printf("Выход: %d.\n", sliceMultiplication(original))
}
func sliceMultiplication(original []int) []int {
	for i, num := range original {
		original[i] = num * 2
	}
	return original
}
