// 5. Напишите функцию InvertMap, которая принимает мапу map[string]int и возвращает новую мапу map[int]string,
// где ключи и значения поменялись местами. Если в исходной мапе есть дублирующиеся значения, вернуть ошибку.
package main

import (
	"fmt"
)

func main() {
	m := map[string]int{"отдых": 2, "респект": 3, "дом": 2}
	m1, err := InvertMap(m)
	if err != nil {
		fmt.Println("Ошибка: ", err)
	} else {
		fmt.Println(m1)
	}

}
func InvertMap(a map[string]int) (map[int]string, error) {
	b := make(map[int]string)
	for k, v := range a {
		if _, ok := b[v]; ok {
			return nil, fmt.Errorf("есть дублирующиеся значения")
		}
		b[v] = k
	}
	return b, nil
}
