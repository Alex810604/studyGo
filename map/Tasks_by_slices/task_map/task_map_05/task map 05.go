// 5. Вложенные мапы (словарь словарей)
//Создайте map[string]map[string]int, где внешний ключ — имя студента, а внутренняя мапа хранит предметы и оценки.
//Напишите функцию для добавления оценки студенту.
//Пример:
//Вход:

// db := map[string]map[string]int{}
// addGrade(db, "Alice", "Math", 90)
// Выход: map[Alice:map[Math:90]]
package main

import "fmt"

func main() {
	db := map[string]map[string]int{}
	//Для накопления оценок
	db2 := map[string]map[string][]int{"Иван": {"Матем": {9}}}
	var name string
	var subject string
	var grade int
	fmt.Print("Введите имя: ")
	fmt.Scan(&name) //Добавляем имя
	fmt.Print("Введите предмет: ")
	fmt.Scan(&subject) //Добавляем предмет
	fmt.Print("Введите оценку: ")
	fmt.Scan(&grade) //Добавляем оценку
	addGrade(db, name, subject, grade)
	addGrade2(db2, name, subject, grade)
	fmt.Println("==========================")
	fmt.Print("Результат по условию задачи: ")
	fmt.Println(db)
	//Список всех оценок
	fmt.Println("-----------------------------------------")
	fmt.Print("Результат перечисления оценок: ")
	fmt.Println(db2)
}

// Функция по условию задачи
func addGrade(db map[string]map[string]int, name, subject string, grade int) {
	if db[name] == nil {
		db[name] = make(map[string]int)
	}
	db[name][subject] = grade
}

// Функция перечисления всех оценок
func addGrade2(db2 map[string]map[string][]int, name, subject string, grade int) {
	if db2[name] == nil {
		db2[name] = make(map[string][]int)
	}
	db2[name][subject] = append(db2[name][subject], grade)
}
