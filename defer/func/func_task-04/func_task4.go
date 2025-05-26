// 4. Удаление дубликатов из слайса
// Напишите функцию RemoveDuplicates, которая принимает слайс строк и возвращает новый слайс без дубликатов.
package main

import "fmt"

func main() {
	n := []string{"дом", "том", "респект", "респект", "респект"}
	fmt.Println(RemoveDuplicates(n))
}
func RemoveDuplicates(d []string) []string {
	m := make(map[string]struct{})
	d1 := make([]string, 0, len(d))
	duplicat := func(str string) {
		if _, ok := m[str]; !ok {
			m[str] = struct{}{}
			d1 = append(d1, str)
		}
	}
	for _, v := range d {
		duplicat(v)
	}
	return d1
}
