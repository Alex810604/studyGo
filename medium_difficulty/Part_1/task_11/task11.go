package main

import (
	"fmt"
	"math"
)

func main() {
	var EuropeanVelocity float64 = 120.4                      // скорость в м/с
	var AmericanVelocity float64 = 130                        // скорость в м/с
	EuropeanVelocity *= 3.6                                   //	скорость в км/ч
	AmericanVelocity *= 2.23694                               //	скорость в миль/ч
	EuropeanVelocity = math.Round(EuropeanVelocity*100) / 100 //	банковское округление до двух чисел
	AmericanVelocity = math.Round(AmericanVelocity*100) / 100 //	банковское округление до двух чисел

	fmt.Println(EuropeanVelocity)
	fmt.Println(AmericanVelocity)
}
