//3. Число четное или нечетное
//Напишите программу, которая проверяет, является ли введенное пользователем число четным или нечетным.

package main

import "fmt"

func main() {
	var name int
	fmt.Print("Введите число для проверки его четности: ")
	_, err := fmt.Scanln(&name)
	if err != nil {
		fmt.Printf("Ошибка: %d.\n", err)
		return
	}
	if name%2 == 0 {
		fmt.Printf("Число %d является четным.\n", name)

	} else {
		fmt.Printf("Число %d является нечётным.\n", name)
	}
}
