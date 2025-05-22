// 17. Найти разницу между двумя массивами
// Даны два массива. Верните элементы, которые есть в одном, но нет в другом.
package main

import "fmt"

func main() {
	num1 := []int{1, 3, 5, 7, 8, 9, 11}
	num2 := []int{1, 3, 5, 8, 8, 10, 11}
	d1 := mapSlice(differenceArrays(num1, num2))
	fmt.Println(d1)
}
func differenceArrays(num1, num2 []int) map[int]int {
	m1 := make(map[int]int)
	m2 := make(map[int]int)
	m3 := make(map[int]int)
	for _, n := range num1 {
		m1[n]++
	}
	for _, n := range num2 {
		m2[n]++
	}
	for i, n := range m1 {
		if m2[i] < n {
			m3[i] = n - m2[i] // {7, 9}
		}
	}
	for i, n := range m2 {
		if m1[i] < n {
			m3[i] = n - m1[i] // к существующей {7, 9} добавили {8, 10}
		}
	}
	return m3
}

func mapSlice(m map[int]int) []int {
	slice1 := make([]int, 0, len(m))
	for n := range m {
		slice1 = append(slice1, n)
	}
	return slice1
}
