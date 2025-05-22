// 20. Проверить, является ли строка панграммой
// Панграмма — это строка, содержащая все буквы алфавита. Проверьте, является ли данная строка панграммой
// (можно игнорировать регистр).
package main

import (
	"fmt"
	"strings"
	"unicode"
)

func main() {
	textRussian := "Йц укен ГШЩЗХ67**?;!ъф  в№;апр ол5 458600 95 ыджэячс 3митЬ+-БЮёяя"
	textEnglish := "QWerTYu23i kopA56/,.,SD Fghj l.//*zxc v bnm"

	//Выбор в строке только букв алфавита
	textRu := construction(textRussian)
	textEn := construction(textEnglish)

	b1, b2 := 'а', 'я'                                // начальная и конечная русские буквы
	d1, d2 := 'a', 'z'                                //	-//-				английские буквы
	english1 := choiceLetters(d1, d2)                 //английский алфавит -> abcdefghijklmnopqrstuvwxyz
	russian1 := choiceLetters(b1, b2)                 //русский алфавит -> абвгдежзийклмнопрстуфхцчшщъыьэюя
	russian2 := string(append([]rune(russian1), 'ё')) //"ё" в таблице unicode стоит после всех русских букв

	//Вывод:
	//для En
	if stringPangram(textEn, english1) {
		fmt.Printf("Строка '%s' является панграммой.\n", textEnglish)
	} else {
		fmt.Printf("Строка '%s' не является панграммой.\n", textEnglish)
	}
	//для Ru
	if stringPangram(textRu, russian2) {
		fmt.Printf("Строка '%s' является панграммой.\n", textRussian)
	} else {
		fmt.Printf("Строка '%s' не является панграммой.\n", textRussian)
	}
}

// Упрощение строки для ее дальнейшего сравнения
func construction(text string) string {
	var a strings.Builder
	for _, n := range text {
		if unicode.IsLetter(n) {
			a.WriteRune(n)
		}
	}
	return strings.ToLower(a.String())
}

// Нахождение интервала букв в алфавите
func choiceLetters(a, b rune) string {
	var alphabet []rune
	for c := a; c <= b; c++ {
		alphabet = append(alphabet, c)
	}
	return string(alphabet) //abcdefghijklmnopqrstuvwxyz или абвгдежзийклмнопрстуфхцчшщъыьэюя
}

// Нахождение строки-панграммы
func stringPangram(text, alphabet string) bool {
	m := make(map[rune]bool)
	lengthAlphabet := len([]rune(alphabet))
	for _, n1 := range text {
		for _, n2 := range alphabet { //алфавит
			if n1 == n2 && !m[n1] {
				m[n1] = true // map из рун
			}
		}
	}
	if len(m) != lengthAlphabet {
		return false
	}
	return true
}
