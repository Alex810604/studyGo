// 10. Обработка ошибок в методах
// Задача: Создайте структуру Wallet с полем Balance.
// Добавьте метод Spend(amount float64) error, который уменьшает баланс, но возвращает ошибку, если денег недостаточно.
package structure

import "errors"

type Wallet struct {
	Balance float64
}

func (w Wallet) Spend(amount float64) error {
	if w.Balance < amount {
		return errors.New("баланс должен быть больше или равен сумме")
	}
	w.Balance -= amount
	return nil
}
