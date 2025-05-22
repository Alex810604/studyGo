// 3. Удаление элемента по индексу
// Напишите функцию, которая удаляет элемент из среза по заданному индексу in-place (без возврата нового среза).
// Пример:
// Вход: slice = []int{10, 20, 30, 40}, index = 1
// Выход: [10 30 40]
package main

import "fmt"

func main() {
	var slice = []int{10, 20, 30, 40}
	var index uint = 1
	copy(slice[index:], slice[index+1:])
	fmt.Println(slice[:len(slice)-1])
}
