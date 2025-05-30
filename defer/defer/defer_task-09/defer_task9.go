// 9. Defer и возврат значения в условном блоке.
// Напишите функцию, где defer изменяет возвращаемое значение, но сам return находится внутри if.
// Покажите, как defer влияет на результат.
package main

import "fmt"

func main() {
	x := 3
	y := 2
	fmt.Print(bar(x, y)) //(4)
}

func bar(a, b int) (result int) {
	result = a + b
	defer func() {
		if result > 30 {
			result += 10 //(3.1)
			fmt.Println("a", result)
		} else {
			result *= 10 //или (3.2)
			fmt.Println("b", result)
		}
	}()
	result += 10 //(1)
	fmt.Println("c", result)
	if result > 30 {
		fmt.Println("d", result+5)
		return result + 5 //(2.1)
	}
	fmt.Println("e", result*2)
	return result * 2 //или (2.2)
}
