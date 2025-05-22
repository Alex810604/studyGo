// Напишите программу, которая принимает строку от пользователя
// и выводит ее в обратном порядке.
package main

import (
	"bufio"
	"fmt"
	"os"
	"unicode/utf8"
)

func main() {
	var stroka string
	fmt.Print("Введите строку, чтобы вывести ее в обратном порядке: ")
	s := bufio.NewScanner(os.Stdin) //создание скана для ввода
	//читаем ввод
	s.Scan()
	stroka = s.Text()
	//fmt.Println("Вы ввели: ", stroka)
	isNaoborot(stroka)

}
func isNaoborot(stroka string) {
	fmt.Print("Строка наоборот: ")  // Выводим заголовок перед циклом, который преобразует текст наоборот
	runess := []rune(stroka)        // переводим строку в руны для корректной обработки Unicode
	size := utf8.RuneLen(runess[3]) // возврат байт
	fmt.Println(size)
	fmt.Println(runess[3])

	dlStr := len(runess) // подсчет символов независимо от веса байтов
	if dlStr == 0 {      // проверка на пустую строку
		fmt.Println("Строка пустая")
	}
	if dlStr > 0 {
		for i := dlStr - 1; i >= 0; i-- {
			g := string(runess[i])
			fmt.Print(g)
		}
		fmt.Println()
	}
}
