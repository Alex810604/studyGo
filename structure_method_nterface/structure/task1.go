//1. Простая структура и метод
//Задача: Создайте структуру Rectangle с полями Width и Height.
//Добавьте метод Area(), который вычисляет площадь прямоугольника. метод - ?

package structure

type Rectangle struct {
	Width  float64
	Height float64
}

func Area(r Rectangle) float64 { // метод или функция
	return r.Width * r.Height

}

//func main() {
//	r := Rectangle{5, 4.36}
//	fmt.Println("Площадь прямоугольника:", r.Area())
//}
