// 4. Встраивание структур
// Задача: Создайте структуру Person с полями Name и Age.
// Затем создайте структуру Employee, встраивающую Person, и добавляющую поле Salary.
package structure

type Person struct {
	Name string
	Age  int
}
type Employee struct {
	Person
	Salary int
}
