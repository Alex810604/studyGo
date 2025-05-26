// 1. Базовый defer.
// Напишите функцию, которая выводит "Start", затем откладывает вывод "Defer", а потом выводит "End".
// Объясните порядок вывода.
package main

import "fmt"

func main() {
	basicDefer()
}

func basicDefer() {
	fmt.Println("Start")
	defer fmt.Println("Defer")
	fmt.Println("End")
}
