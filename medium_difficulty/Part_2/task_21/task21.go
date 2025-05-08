package main

import (
	"fmt"
	"math"
)

func main() {
	length := 35.00
	var radius float64
	var R *float64
	R = &radius
	*R = length / (2 * math.Pi)
	AreaCircle := math.Pi * math.Pow(*R, 2)
	fmt.Printf("Pадиус окружности: %.2f\nПлощадь круга: %.2f\n", *R, AreaCircle)
}
