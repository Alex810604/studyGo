// 5. Передача среза по указателю
// Напишите функцию, которая принимает указатель на срез строк и добавляет строку "modified" в конец.
// Пример:
// Вход: &[]string{"a", "b"}
// Выход: ["a" "b" "modified"]
package main

import "fmt"

func main() {
	slice := []string{"a", "b"}
	fmt.Printf("Вход: %v.\n", slice)
	takesPointer(&slice)
	fmt.Printf("Выход: %v.\n", slice)
}
func takesPointer(p *[]string) {
	*p = append(*p, "modified")

}
