// 8. Подсчёт среднего значения в слайсе.
// Напишите функцию Average, которая принимает слайс чисел и возвращает их среднее значение.
// Если слайс пуст, возвращает 0.
package main

import "fmt"

func main() {
	numbers := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}

	// Вариант 1
	result1 := Average1(numbers, func(sum, number int) int {
		return sum + number
	})
	fmt.Println(result1)
	// вариант 2
	result2 := Average2(numbers)
	fmt.Println(result2)
	// вариант 3
	result3 := Average3(numbers)
	fmt.Println(result3)
}

// вариант 1
func Average1(numbers []int, d func(int, int) int) float64 {
	sum := 0
	for _, number := range numbers {
		sum = d(sum, number)
	}
	return float64(sum) / float64(len(numbers))
}

// вариант 2
func Average2(numbers []int) float64 {
	sum := 0
	d := func(num int) { sum += num }

	for _, number := range numbers {
		d(number)
	}
	return float64(sum) / float64(len(numbers))
}

// вариант 3
func Average3(numbers []int) float64 {
	sum := 0
	for _, number := range numbers {
		sum += number
	}
	return float64(sum) / float64(len(numbers))
}
