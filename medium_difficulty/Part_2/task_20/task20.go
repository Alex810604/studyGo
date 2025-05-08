package main

import "fmt"

func main() {
	var A *int
	var B int = 3
	A = &B
	fmt.Println(*A)
	*A = 10
	fmt.Println(B)
}
