// 9. Сгруппировать слова по длине
// Дан список слов. Сгруппируйте их в хешмапу, где ключ — длина слова, а значение — список слов такой длины.
package main

import "fmt"

func main() {
	words := []string{"дед", "дом", "авто", "имя", "книга", "жизнь"}
	m := make(map[int][]string) //int	string
	d := 0
	for _, v := range words {
		d = len([]rune(v))
		m[d] = append(m[d], v)
	}
	fmt.Println(m)

}
