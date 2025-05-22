// 9. Вставка элемента по индексу.
// Напишите функцию, которая вставляет элемент в срез по указанному индексу с сохранением исходного capacity.
// Пример:
// Вход: slice = []int{10, 20, 30}, index = 1, value = 99
// Выход: [10 99 20 30]
package main

import "fmt"

func main() {
	slice := []int{10, 20, 30}
	value := 99
	index := 1
	if cap(slice) >= len(slice)+1 {
		nSlice := slice[:len(slice)+1]
		copy(nSlice, slice[:index])
		nSlice[index] = value
		copy(nSlice[index+1:], slice[index:])
		fmt.Println(nSlice)
	} else {
		nSlice := make([]int, len(slice)+1)
		copy(nSlice, slice[:index])
		nSlice[index] = value
		copy(nSlice[index+1:], slice[index:])
		fmt.Println(nSlice)
	}
}
