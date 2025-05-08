// 1. Напишите функцию Swap(a, b *int), которая меняет местами значения двух целых чисел, используя указатели.
package main

import (
	"fmt"
)

func main() {
	var x, y int
	fmt.Printf("Введите значение x: ")
	fmt.Scan(&x)
	fmt.Printf("Введите значение y: ")
	fmt.Scan(&y)
	Swap(&x, &y)
	//fmt.Println(a, b)
	fmt.Printf("Изменненные местами значения 'x' и 'y' равны: %d и %d\n", x, y)
}

func Swap(a, b *int) {
	*a, *b = *b, *a
}
