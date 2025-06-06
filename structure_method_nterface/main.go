package main

import (
	"awesomeProject/structure_method_nterface/structure"
	"fmt"
)

func main() {
	//задача 1
	//r := structure.Rectangle{5.2, 4.36}
	//fmt.Printf("Площадь прямоугольника: %.2f\n", structure.Area(r))

	//задача 2
	//c := structure.Circle{Radius: 5}
	//r := structure.Rectangle{Width: 5.2, Height: 4.36}
	//structure.Figure(&r)
	//structure.Figure(&c)

	//задача 3
	//d := structure.Counter{4}
	//fmt.Println(d.GetValue())
	//d.Increment()
	//fmt.Println(d.GetValue())
	//d.Increment()
	//fmt.Println(d.GetValue())

	//задача 5
	//b := structure.Book{"Title", "Author"}
	//fmt.Println(b)
	//structure.PrintBook(b)

	//задача 6
	//structure.PrintType(42)
	//structure.PrintType("Привет")
	//structure.PrintType(true)
	//structure.PrintType(3.14)
	//structure.PrintType(nil)
	//structure.PrintType([]int{1, 2})
	//structure.PrintType([]string{"kf"})

	//задача 7
	//students := []structure.Student{	//создаём слайс студентов.
	//	{"Алекс", 85},
	//	{"Мария", 90},
	//	{"Иван", 78},
	//	{"Елена", 95},
	//}
	//sort.Sort(structure.ByGrade(students)) // Сортируем
	//fmt.Println(students)
	//for _, s := range students {
	//	fmt.Println(s.Name, "-", s.Grade)
	//}

	//задача 8
	//var _ structure.Database = &structure.MockDB{} // проверка - удовлетворяет ли тип интерфейсу
	//d := structure.MockDB{}
	//d.Save("Данные 1")
	//d.Save("Данные 2")
	//
	//// проверка на пустые строки
	//d.Save("")
	//
	//fmt.Println(d.Read())

	//задача 9
	//d := structure.Dog{}
	//c := structure.Cat{}
	//structure.MakeSound(structure.Dog{})
	//structure.MakeSound(structure.Cat{})

	//задача 10

	//задача 11
	////var rect structure.ExpandShape = structure.Rectangle{5, 10}
	////var circle structure.ExpandShape = structure.Circle{3}
	//rect := structure.Rectangle{Width: 5, Height: 10}
	//circle := structure.Circle{Radius: 3}
	//structure.PrintExpandShape(rect)
	//structure.PrintExpandShape(circle)

	//задача 12
	//us := structure.NewUser("admin", "567")
	//fmt.Println("Login", us.GetLogin())
	//fmt.Println("Password", us.GetPassword())

	//задача 13
	//result, err := structure.Divide(10, 3)
	//if err != nil {
	//	fmt.Println(err)
	//} else {
	//	fmt.Printf("Результат: %.2f\n", result)
	//}

	//задача 14
	//n := structure.Calculator{}
	//fmt.Println(n.Sum(1, 2, 3, 4, 5))
	//numbers := []int{10, 20, 30}
	//fmt.Println(n.Sum(numbers...))
	//fmt.Println(n.Sum())

	//задача 15
	//original := &structure.Document{Content: "Оригинальный текст"}
	//clone := original.Clone().(*structure.Document) // Клонируем объект
	//fmt.Println("Оригинал:", original.Content)
	//fmt.Println("Копия:", clone.Content)
	//clone.Content = "Модифицированные данные"
	//fmt.Println("Оригинал:", original.Content)
	//fmt.Println("Копия:", clone.Content)

	//задача 16
	//event := &structure.Event{
	//	Title: "Встреча с клиентом",
	//	Time:  time.Now().Add(-5 * time.Minute),
	//}
	//fmt.Println(*event)
	//if event.IsAfter(time.Now()) {
	//	fmt.Printf("Событие '%s' еще не наступило\n", event.Title)
	//} else {
	//	fmt.Printf("Событие '%s' уже прошло\n", event.Title)
	//}

	//задача 17
	//products := []structure.Product{
	//	{"авто", 500, 4.5},
	//	{"книга", 120, 5},
	//	{"стол", 400, 4.8},
	//}
	//compator := structure.Comparator{
	//	Product: products,
	//	LessFunc: func(p1, p2 structure.Product) bool {
	//		return p1.Name < p2.Name
	//	},
	//}
	//sort.Sort(compator)
	//fmt.Println(compator.Product)

	//задача 18
	transport := []structure.Transport{structure.Car{}, structure.Bicycle{}, structure.Plane{}}
	structure.StartRace(transport)

	// с добавлением полей в структурах
	fmt.Printf("=======================\n")

	transport1 := []structure.Transport1{
		structure.Car1{"Шкода"},
		structure.Bicycle1{"Stels"},
		structure.Plane1{"SU-123"},
	}
	structure.StartRace1(transport1)
}
