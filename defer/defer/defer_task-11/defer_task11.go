//Скопируйте код в файл не запуская и попробуйте определить что выведет программа
//В решении напиши комментарии и объяснение почему так работает
//===============================================================================
// Что выведет программа и почему?

package main

import "fmt"

type Account struct {
	Balance int
}

func main() {
	initialBalance := 1000
	account := &Account{Balance: initialBalance}

	//defer вычисляет аргументы на момент объявления defer, которые фиксируются, но выполняется функция позже.
	defer printBalance("Изначальный баланс", account.Balance) // 1000  (3)
	defer printBalance("Текущий баланс", account.Balance)     // 1000  (2)

	//в defer передается копия указателя account с текущим значением по этому указателю, которое может изменяться.
	defer printAccountBalance("Указатель на баланс", account) //1300  (1)	account - указатель

	account.Balance += 500      //1500
	updateBalance(account, 200) //1300	account - указатель, принимает текущее значение параметра

	// Здесь account принимает новый указатель с его балансом.
	account = &Account{Balance: 300} //300
	//defer работает со старым указателем на момент его объявления и его баланс принимается на момент вычисления.
}

func updateBalance(acc *Account, amount int) {
	acc.Balance -= amount
}

func printBalance(label string, amount int) {
	fmt.Printf("%s: %d\n", label, amount)
}

func printAccountBalance(label string, acc *Account) {
	fmt.Printf("%s: %d\n", label, acc.Balance)
}
