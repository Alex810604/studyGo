//2. Сложение двух чисел
//Создайте программу, которая запрашивает у пользователя два числа и выводит их сумму.

package main

import "fmt"

func main() {
	var a, b int
	fmt.Print("введите первое число: ")
	fmt.Scan(&a)
	fmt.Print("введите второе число: ")
	fmt.Scan(&b)
	fmt.Printf("сумма чисел %d и %d = %d \n", a, b, a+b)
}
