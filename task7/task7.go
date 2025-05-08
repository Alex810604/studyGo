// Напишите функцию, которая проверяет, является ли строка палиндромом.
package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {
	var stroka string
	s := bufio.NewScanner(os.Stdin)
	fmt.Print("Введите строку: ")
	s.Scan()
	stroka = s.Text()
	//fmt.Println("Вы ввели: ", stroka)
	glagol := "не"
	if strPalindrome(stroka) {
		glagol = ""
	}
	fmt.Printf("Строка '%s' %s является палиндромом.\n", stroka, glagol)
}
func strPalindrome(stroka string) bool {
	str := strings.ToLower(stroka) // Преобразует все символы в нижний регистр

	// Удаление всех пробелов
	strokaBezprobelow := strings.ReplaceAll(str, " ", "") // Замена старого символа на новый по всей строке (удаление пробелов)
	//fmt.Println("Текст без пробелов: ", strokaBezprobelow)
	runess := []rune(strokaBezprobelow)              // переводим строку в руны для корректной обработки Unicode
	dlStr := len(runess)                             // подсчет символов независимо от веса байтов
	for i, j := 0, dlStr-1; i < j; i, j = i+1, j-1 { //                                или for i := 0; i < dlStr/2; i++ {
		if runess[i] != runess[j] { // Если найдена хотя бы одна несовпадающая пара	    //       if runess[i] != runess[dlStr-1-i]
			return false // Немедленный возврат false - это НЕ палиндром
		}
	}
	// Если все пары символов совпали
	return true // Возврат true - это палиндром
}
