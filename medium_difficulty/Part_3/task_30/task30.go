package main

import "fmt"

func main() {
	a := []int{1, 2, 3, 4, 5, 6, 7}
	b := make([]int, 5)
	copy(b, a)
	a = a[5:]
	//var b []int

	fmt.Printf("Слайс с выходными днями: %d\nСлайс с рабочими днями: %d\n", a, b)
}
