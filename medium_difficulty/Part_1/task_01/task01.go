package main

import (
	"fmt"
	"strconv"
)

func main() {
	var a string = "104"
	var b int = 35
	//a1 := strconv.Atoi(a) //  string → int
	a1, err := strconv.Atoi(a)
	if err != nil {
		fmt.Println(err)
		return
	}
	b1 := strconv.Itoa(b) //  int → string
	fmt.Println(a1, b1)

}
