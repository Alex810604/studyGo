// 7. Поиск уникальных элементов в слайсе.
// Напишите функцию FindUnique, которая принимает слайс целых чисел и возвращает новый слайс,
// содержащий только уникальные элементы (встречающиеся 1 раз).
package main

import "fmt"

func main() {
	numbers := []int{1, 2, 2, 4, 4, 6, 7, 9, 9, 10}
	// вариант 1
	result1 := FindUnique1(numbers, func(i int) bool { return i == 1 })
	fmt.Println(result1)
	// вариант 2
	result2 := FindUnique2(numbers)
	fmt.Println(result2)
	// вариант 3
	result3 := FindUnique3(numbers)
	fmt.Println(result3)
}

// вариант 1
func FindUnique1(numbers []int, d func(i int) bool) []int {
	var slice []int
	m := make(map[int]int)
	for _, number := range numbers {
		m[number]++
	}
	for k, number := range m {
		if d(number) {
			slice = append(slice, k)
		}
	}
	return slice
}

// вариант 2
func FindUnique2(numbers []int) []int {
	var slice []int
	m := make(map[int]int)
	for _, number := range numbers {
		m[number]++
	}
	d := func(num int) bool {
		return num == 1
	}
	for k, number := range m {
		if d(number) {
			slice = append(slice, k)
		}
	}
	return slice
}

// вариант 3
func FindUnique3(numbers []int) []int {
	var slice []int
	m := make(map[int]int)
	for _, number := range numbers {
		m[number]++
	}
	for _, number := range numbers {
		if m[number] == 1 {
			slice = append(slice, number)
		}
	}
	return slice
}
