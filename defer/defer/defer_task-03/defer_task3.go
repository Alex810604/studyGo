// 3. Defer и возвращаемые значения.
// Напишите функцию, которая использует defer для изменения возвращаемого значения. Объясните результат.
package main

import "fmt"

func main() {

	fmt.Println(bar())
	fmt.Println(double(5))

}
func double(x int) (result int) {
	defer func() { result *= 3 }() // без defer результат функции равен 0 -> 0*3=0
	return x                       // 15, без defer return выведет 5, так как присвоит новое значение -> x
}

func bar() (result int) {
	result = 40
	defer func() { result += 5 }()
	result *= 2
	return result //85
}
