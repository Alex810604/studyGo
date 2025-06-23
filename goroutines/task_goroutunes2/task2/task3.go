// 3. Конфигурация с горячей перезагрузкой
// Разработайте систему управления конфигурацией, которая позволяет:
// Часто читать параметры конфигурации (много горутин)
// Редко полностью перезагружать конфигурацию (одна горутина)
// Используйте sync.RWMutex.
package goroutine2

import (
	"fmt"
	"sync"
)

type Config struct {
	params map[string]interface{}
	mu     sync.RWMutex
}

func NewConfig(newParams map[string]interface{}) *Config {
	return &Config{params: newParams}
}

// PrintConfig выводит текущую конфигурацию
func (c *Config) PrintConfig() {
	fmt.Println("Текущая конфигурация:")
	for key, value := range c.params {
		fmt.Printf("  %s: %v\n", key, value)
	}
}

// Get возвращает значение параметра конфигурации
func (c *Config) Get(key string) interface{} {
	c.mu.RLock()
	defer c.mu.RUnlock()
	return c.params[key]
}

// Reboot полностью перезагружает конфигурацию
func (c *Config) Reboot(newParams map[string]interface{}) {
	c.mu.Lock()
	defer c.mu.Unlock()
	fmt.Println("Конфигурация перезагружена!")
	c.params = newParams
	c.PrintConfig()
}

func (c *Config) PrintSafe(id int, name, version interface{}) {
	c.mu.RLock()
	defer c.mu.RUnlock()
	fmt.Printf("Горутина %d: name= %v, version = %v\n", id, name, version)
}
