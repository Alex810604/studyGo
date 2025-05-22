// 4. Найти первый неповторяющийся символ
// Дана строка. Найдите первый символ, который встречается только один раз.
package main

import "fmt"

func main() {
	str := "павел левак в поле ковал"
	//n := ""
	m := make(map[string]int)
	for _, v := range str {
		m[string(v)]++
	}
	for _, v := range str {
		if m[string(v)] == 1 {
			fmt.Println(string(v))
			return
		}
	}
	fmt.Printf("В строке '%v' нет универсальных символов: %v.\n", str, m)

}
