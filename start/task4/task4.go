// 4. Факториал числа
//Реализуйте функцию, которая вычисляет факториал заданного числа.

package main

import (
	"fmt"
	"strconv"
)

func main() {
	var chislo string
	var factorial uint64 = 1
	fmt.Print("Введите число для определения факториала: ")
	fmt.Scan(&chislo)               //ввод данных
	num, er := strconv.Atoi(chislo) //преобразование строки "chislo" в число num (функция - strconv.Atoi)
	if er != nil {                  // выполнение условия на отсутствие ошибки, nil - отсутствие ошибки
		fmt.Printf("Ошибка ввода: '%v' не целое число\n", chislo)
		return
	}
	if num < 0 {
		fmt.Println("Факториал должен быть неотрицательным")
		return
	}
	if num == 1 || num == 0 {
		fmt.Printf("Факториал %d равен 1\n", num)
		return
	}
	for i := 2; i <= num; i++ {
		factorial *= uint64(i)

	}
	fmt.Printf("Факториал числа %d = %d\n", num, factorial)
}
