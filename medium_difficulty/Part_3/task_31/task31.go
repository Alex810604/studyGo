package main

import "fmt"

func main() {
	a := []int{6, 7}
	b := []int{1, 2, 3, 4, 5}
	fmt.Println(append(a, b...))

	//или без переаллокации (сар=7)
	c := make([]int, 0, len(a)+len(b)) // сар=7
	c = append(append(c, a...), b...)
	fmt.Println(c)
}
