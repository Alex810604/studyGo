// 19. Найти слово с максимальной частотой
// Дан текст. Найдите слово, которое встречается чаще всего (регистр не учитывать).
package main

import (
	"fmt"
	"strings"
	"unicode"
)

func main() {
	text := "Привет! Дом построен? - спросил   он. Да  ,   привет, дом построен и  ПОСТРОЕН  гараж."
	var a strings.Builder                     // тип строки
	for _, n := range strings.ToLower(text) { // сразу перевод в нижний регистр, значения в рунах
		if unicode.IsLetter(n) || unicode.IsDigit(n) || unicode.IsSpace(n) { // буквы, числа, пробелы
			a.WriteRune(n) // вписывает значения 'n' в рунах в строку 'а'
		}
	}
	//d := a.String() // Вывод: "привет дом построен  спросил   он да     привет дом построен"
	d1 := strings.Fields(a.String()) // Вывод слайса: [привет дом построен спросил он да привет дом построен]
	//str := strings.Join(d1, " ")	//Вывод обратно в строку: "привет дом построен спросил он да привет дом построен и построен гараж"
	m := make(map[string]int)
	for _, n := range d1 {
		m[n]++
	}
	key, total := "", 0
	for i, v := range m {
		if v > total {
			key, total = i, v
		}
	}
	fmt.Printf("Чаще всего встречается слово '%s' -> %d раз(а).\n", key, total)
}

//fmt.Println(removePunctuation(text))
//func removePunctuation(text string) string {
//	var result []rune
//	for _, char := range text {
//		if unicode.IsLetter(char) || unicode.IsDigit(char) || unicode.IsSpace(char) {
//			result = append(result, char)
//		}
//	}
//	return string(result)
//}
