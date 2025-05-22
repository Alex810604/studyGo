// 1.Напишите функцию, которая создает срез целых чисел от 1 до n,
// затем добавляет к нему числа от n+1 до 2n с помощью append.
// Выведите исходный и измененный срез.
// a := append(a,
// Пример:
// Вход: n = 3
// Выход:
// Исходный: [1 2 3]
// После append: [1 2 3 4 5 6]
package main

import "fmt"

func main() {
	var n int
	fmt.Print("Введите число n: ")
	fmt.Scan(&n)
	original, sliceModified := sliceIntegers(n)
	fmt.Printf("Исходный: %d.\nИзмененный срез:  %d.\n", original, sliceModified)
}
func sliceIntegers(n int) ([]int, []int) {
	original := make([]int, n)
	for i := 0; i < n; i++ {
		original[i] = i + 1
	}

	// создание копии среза
	sliceModified := append(original[:0:0], original...)

	for i := n + 1; i <= 2*n; i++ {
		sliceModified = append(sliceModified, i)
	}
	return original, sliceModified
}
