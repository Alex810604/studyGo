// 11. Интерфейс для геометрических фигур с периметром
// Задача:
// Расширьте интерфейс Shape из задачи 2, добавив метод Perimeter() float64. Реализуйте его для Rectangle и Circle.
package structure

import (
	"fmt"
	"math"
)

type ExpandShape interface {
	Shape
	Perimeter() float64
}

func (r Rectangle) Perimeter() float64 {
	return (r.Width + r.Height) * 2
}

func (c Circle) Perimeter() float64 {
	return 2 * c.Radius * math.Pi
}
func PrintExpandShape(e ExpandShape) {
	fmt.Printf("Площадь: %.2f. Периметр (длина оружности): %.2f.\n", e.Area(), e.Perimeter())
}
