// 1. Подсчёт частоты элементов в слайсе
// Напишите функцию CountFrequency, которая принимает слайс строк и возвращает мапу,
// где ключи — элементы слайса, а значения — количество их вхождений.
package main

import "fmt"

func main() {
	n := []string{"дом", "том", "респект", "респект", "респект"}
	fmt.Println(CountFrequency(n))

}
func CountFrequency(words []string) map[string]int {
	m := make(map[string]int)
	count := func(str string) { m[str]++ }
	for _, v := range words {
		count(v)
	}
	return m
}
