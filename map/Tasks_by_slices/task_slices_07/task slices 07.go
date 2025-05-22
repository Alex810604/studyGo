// 7. Фильтрация in-place.
// Напишите функцию, которая удаляет из среза все четные числа in-place (без выделения нового среза).
// Пример:
// Вход: []int{1, 2, 3, 4, 5}
// Выход: [1 3 5]
package main

import "fmt"

func main() {
	slice := []int{1, 2, 6, 4, 5}
	for i := 0; i < len(slice); i++ {
		if slice[i]%2 != 0 {
			continue
		}
		if slice[i]%2 == 0 {
			copy(slice[i:], slice[i+1:])
			slice = slice[:len(slice)-1]
			i--
			//fmt.Println(slice)
		}
	}
	fmt.Println(slice)
}
