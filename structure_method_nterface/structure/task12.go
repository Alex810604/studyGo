// 12. Конструкторы и приватные поля
// Задача: Создайте структуру User с приватными полями login и password.
// Напишите конструктор NewUser(login, password string) *User и метод GetLogin() string, который возвращает логин.
package structure

type User struct {
	login    string
	password string
}

func (u *User) GetLogin() string {
	return u.login
}
func (u *User) GetPassword() string {
	return u.password
}
func NewUser(login, password string) *User {
	return &User{login: login, password: password}
}
