// 5. Банковский счет с RW Mutex
// Создайте структуру банковского счета, где:
// Проверка баланса происходит очень часто
// Изменение баланса (пополнение/списание) происходит редко
// Оптимизируйте с помощью sync.RWMutex.
package goroutine2

import (
	"fmt"
	"sync"
	"time"
)

type BankАccount struct {
	Balance int
	rwmu    sync.RWMutex
}

func (b *BankАccount) ViewBalance() int {
	b.rwmu.RLock()
	defer b.rwmu.RUnlock()
	time.Sleep(5 * time.Millisecond)
	return b.Balance
}

func (b *BankАccount) CreditingBalance(value int) {
	b.rwmu.Lock()
	defer b.rwmu.Unlock()
	time.Sleep(6 * time.Millisecond)
	b.Balance += value
	fmt.Println("Ваш баланс увеличен: ", b.Balance)

}

func (b *BankАccount) ReduceBalance(value int) {
	b.rwmu.Lock()
	defer b.rwmu.Unlock()
	time.Sleep(4 * time.Millisecond)
	if b.Balance < value {
		fmt.Printf("Ошибка списания: недостаточно средств (баланс: %d, запрошено: %d)\n", b.Balance, value)
	} else {
		b.Balance -= value
		fmt.Println("Баланс списан и составляет: ", b.Balance)
	}
}
