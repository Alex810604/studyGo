// 2. Фильтрация слайса по условию.
// Напишите функцию FilterEven, которая принимает слайс целых чисел и возвращает новый слайс, содержащий только чётные числа.
package main

import "fmt"

func main() {
	n := []int{1, 2, 3, 4, 5}
	fmt.Println(FilterEven(n))

}
func FilterEven(nums []int) []int {
	var nums1 []int
	count := func(d int) bool { return d%2 == 0 }
	for _, v := range nums {
		if count(v) {
			nums1 = append(nums1, v)
		}
	}
	return nums1
}
