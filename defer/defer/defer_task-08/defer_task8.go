// 8. Defer и цепочка вызовов.
// Напишите функцию с несколькими defer, где один defer вызывает другую функцию, которая тоже содержит defer.
// Покажите порядок выполнения.
package main

import "fmt"

func main() {
	a := 1
	defer func() {
		a *= 4 //(6) ->120
		fmt.Println("a=", a)
	}()
	a += 2 //(1) ->3
	fmt.Println("a=", a)
	defer func() {
		a *= 3 //(3) ->12
		fmt.Println("a=", a)
		add(&a)
	}()
	a++ //(2) ->4
	fmt.Println("a=", a)
}
func add(a *int) {
	defer func() { *a += 10; fmt.Println("a=", *a) }() //(5)-> a= 30
	*a += 8                                            //(4)-> a= 20
	fmt.Println("a=", *a)

}
