// 6. Перебор мапы
// Дана мапа colors := map[string]string{"red": "#FF0000", "green": "#00FF00", "blue": "#0000FF"}.
// Выведите все ключи и значения в формате "red: #FF0000".
package main

import "fmt"

func main() {
	colors := map[string]string{"red": "#FF0000", "green": "#00FF00", "blue": "#0000FF"}
	if colors["red"] == "#FF0000" {
		colors = map[string]string{"red": "#FF0000"} // Создаём новый map только с "red"
	}
	fmt.Println(colors)
	//или
	colors = map[string]string{"red": "#FF0000", "green": "#00FF00", "blue": "#0000FF"}
	if value, exists := colors["red"]; exists && value == "#FF0000" {
		colors = map[string]string{"red": value}
	}
	fmt.Println(colors)
}
