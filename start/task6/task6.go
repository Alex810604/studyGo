// Создайте программу, которая находит максимальное число в массиве целых чисел.
package main

import "fmt"

func main() {
	var n int64
	var numbers int // количество чисел
	var num []int64
	fmt.Print("Введите количество чисел в массиве: ")
	fmt.Scanln(&numbers)
	for i := 0; i < numbers; i++ {
		fmt.Printf("Введите число №%d: ", i+1)
		fmt.Scan(&n)         //создаем скан и вводим числа в массиве
		num = append(num, n) // добавляем сканированные числа в срез num
	}
	//fmt.Println("Введенные числа:", num)

	// цикл нахождения максимального числа
	max := num[0]            //первое число принимаем за максимум
	for _, dd := range num { //перебираем все числа в срезе num
		if dd > max {
			max = dd
		}
	}
	fmt.Printf("Максимальное число: %d", max)
	fmt.Println()
}
