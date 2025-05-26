// 9. Поиск дубликатов в слайсе
// Дан слайс nums := []int{1, 2, 3, 2, 4, 5, 4}. Проверьте, есть ли в нём дубликаты, используя мапу.
package main

import "fmt"

func main() {
	nums := []int{1, 3, 2, 4, 5, 4}
	m := make(map[int]int)
	var num []int
	for _, r := range nums {
		m[r]++
	}
	fmt.Println(m)
	for key, value := range m {
		if value > 1 {
			num = append(num, key)
		}
	}
	if len(num) == 0 {
		fmt.Println("дубликатов нет")
	}
	if len(num) > 0 {
		fmt.Printf("дубликаты: %d", num)
	}

	//

	//	fmt.Printf("Ключ-символ %d встречается %d раз\n", k, v)

}
