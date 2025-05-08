// Напишите функцию IsNil(val *int) bool, которая проверяет, является ли переданный указатель nil.
package main

import "fmt"

func main() {
	var ptr *int
	var num int
	fmt.Println(IsNil(ptr))
	ptr = &num
	fmt.Println(IsNil(ptr))
}
func IsNil(val *int) bool {
	return val == nil
}
