package main

import (
	"fmt"
)

func contains(a []string, x string) bool {
	str := ""
	for _, str = range a {
		if str == x {
			return true
		}
	}
	return false
}
func getMax(nums ...int) int {
	if len(nums) == 0 {
		return 0
	}
	max := nums[0]
	for _, num := range nums[1:] {
		if num > max {
			max = num
		}
	}
	return max
}

func main() {
	var x string = "Иван"
	var a []string = []string{"Иван", "Анна"}
	fmt.Println(getMax(1, 5, 56, 88))
	d := contains(a, x)
	fmt.Println(d)
}
