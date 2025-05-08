// Реализуйте программу, которая вычисляет сумму всех элементов в массиве.
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
	sumMassiva(num)
}

func sumMassiva(num []int64) {
	var sum int64 = 0
	for _, value := range num {
		sum += value
	}
	fmt.Printf("сумма всех элементов в массиве %d равна %d.\n", num, sum)
}
