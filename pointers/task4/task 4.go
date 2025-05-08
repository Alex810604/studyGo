// Напишите функцию CreatePointer(x int) *int, которая возвращает указатель на переданное число.
package main

import "fmt"

func main() {
	var a int = 100
	ptr := CreatePointer(a)
	fmt.Println(*ptr) // разыменовываем (получаем) значение 'a'
	//и выведем саму ячейку памяти
	fmt.Println(ptr)
}
func CreatePointer(x int) *int {
	//x += 12
	return &x
	// Значение return &x компилятор Go размещает x в куче, потому что возвращается указатель,
	// и значение x должно сохраниться после вызова функции.
}
