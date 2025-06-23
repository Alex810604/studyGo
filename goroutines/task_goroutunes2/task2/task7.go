// 7. Кэш статистики с редкими обновлениями
// Создайте кэш статистических данных, где:
// Данные обновляются раз в минуту (одна горутина)
// Данные читаются сотни раз в секунду (много горутин)
// Примените sync.RWMutex.
package goroutine2

import (
	"fmt"
	"sync"
)

type Statistics struct {
	cache map[string]interface{}
	rwmu  sync.RWMutex
}

func NewStatistics(newCache map[string]interface{}) *Statistics {
	return &Statistics{cache: newCache}
}

func (s *Statistics) GetStatistics(key string) interface{} {
	s.rwmu.RLock()
	defer s.rwmu.RUnlock()
	return s.cache[key]
}
func (s *Statistics) Update(newData map[string]interface{}) {
	s.rwmu.Lock()
	defer s.rwmu.Unlock()
	s.cache = newData
	fmt.Printf("Новые данные:\n- язык программирования: %v;\n- задача №: %v.\n",
		s.cache["язык программирования"],
		s.cache["задача №"])
}
