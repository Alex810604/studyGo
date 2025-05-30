// 6. Defer и изменение именованного возвращаемого значения (в Go можно типу возвращаемого значения присвоить имя).
// Напишите функцию, где defer изменяет именованное возвращаемое значение после return.
package main

import "fmt"

func main() {
	n := 5
	d := changeMeaning(n)
	fmt.Println("финиш =", d)
}
func changeMeaning(a int) (result int) {
	defer func() { result *= 3; fmt.Println("defer =", result) }()
	result = a * 10
	fmt.Println("С начало result =", result)
	return result
}
