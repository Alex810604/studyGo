package main

import (
	goroutine2 "awesomeProject/goroutines/task_goroutunes2/task2"
	"fmt"
	"sync"
	"time"
)

var wg sync.WaitGroup

var mu sync.Mutex

func main() {
	//Задача 1.
	{
		//words := []string{"Что значит panic...", "Расчет инкремента...", "Как сохранить данные..."}
		////Писатели
		//for i, word := range words {
		//	wg.Add(1)
		//	go goroutine2.Writer(&wg, i, word)
		//	wg.Wait() //Ждем завершения текущего писателя
		//	// Запускаем читателей после каждой записи
		//	for j := 1; j <= 5; j++ {
		//		wg.Add(1)
		//		go goroutine2.Reader(&wg, j)
		//	}
		//	wg.Wait() // Ждём завершения всех читателей
		//}
	}
	//Задача 2. Атомарный счетчик
	{
		//var counter goroutine2.CounterAtomic
		//for i := 0; i < 100; i++ {
		//	wg.Add(1)
		//	go func() {
		//		defer wg.Done()
		//		counter.Increment()
		//	}()
		//}
		//wg.Wait()
		//fmt.Println(counter.CurrentValue())
	}
	//Задача 3.
	{
		//config := goroutine2.NewConfig(map[string]interface{}{ // Инициализация конфигурации
		//	"name":    "GO",
		//	"version": 1.23,
		//})
		//// Выводим начальную конфигурацию
		//config.PrintConfig()
		//
		//for i := 0; i < 500; i++ { // Много горутин читают конфигурацию
		//	wg.Add(1)
		//	go func(id int) {
		//		defer wg.Done()
		//		time.Sleep(100 * time.Microsecond) //для наглядной видимости перезагрузки
		//		name := config.Get("name")
		//		version := config.Get("version")
		//		config.PrintSafe(id, name, version)
		//	}(i)
		//}
		//
		//// Одна горутина перезагружает конфигурацию
		//wg.Add(1)
		//go func() {
		//	defer wg.Done()
		//	mu.Lock()
		//	defer mu.Unlock()
		//	config.Reboot(map[string]interface{}{
		//		"name":    "Goland",
		//		"version": 1.24,
		//	})
		//}()
		//wg.Wait()
	}
	//Задача 4.
	{
		//flag := goroutine2.Flag{}
		//// Запускаем 10 горутин, которые устанавливают флаг
		//for i := 0; i < 10; i++ {
		//	wg.Add(1)
		//	go func(id int) {
		//		defer wg.Done()
		//		flag.Set(id%2 == 0)
		//		fmt.Printf("Горутина %d установила %v\n", id, flag.Get())
		//	}(i)
		//}
		//wg.Wait()
		//fmt.Println("Финальное значение:", flag.Get())
	}
	//Задача 5.
	{
		//client := &goroutine2.BankАccount{Balance: 500}
		//value := 400
		//for i := 0; i < 500; i++ {
		//	wg.Add(1)
		//	go func(id int) {
		//		defer wg.Done()
		//		time.Sleep(10 * time.Millisecond)
		//		fmt.Printf("Горутина %d: ваш баланс: %d\n", id, client.ViewBalance())
		//
		//	}(i)
		//}
		//for j := 0; j < 5; j++ {
		//	wg.Add(1)
		//	go func() {
		//		defer wg.Done()
		//		time.Sleep(15 * time.Millisecond)
		//		client.CreditingBalance(value)
		//	}()
		//	wg.Add(1)
		//	go func() {
		//		defer wg.Done()
		//		time.Sleep(15 * time.Millisecond)
		//		client.ReduceBalance(value)
		//	}()
		//}
		//wg.Wait()
		//fmt.Println("Финальный баланс:", client.ViewBalance())
	}
	//Задача 6.
	{
		//generator := goroutine2.NewAtomicGenerator()
		//for i := 0; i < 10; i++ {
		//	wg.Add(1)
		//	go func(in int) {
		//		defer wg.Done()
		//		genID := generator.ReturnUniqueNumbers()
		//		fmt.Println("Горутина", in, "получила число", genID)
		//	}(i)
		//}
		//wg.Wait()
	}
	//Задача 7.
	{
		//stats := goroutine2.NewStatistics(map[string]interface{}{
		//	"язык программирования": "Go",
		//	"задача №":              17,
		//})
		//
		//// Обновление кэша каждые 10 секунд
		//go func() {
		//	ticker := time.NewTicker(10 * time.Second)
		//	defer ticker.Stop()
		//
		//	languages := []string{"Java", "Python", "Go", "C++"}
		//	tasks := []int{5, 4, 6}
		//	i := 0
		//
		//	for t := range ticker.C {
		//		newCache := map[string]interface{}{
		//			"язык программирования": languages[i%len(languages)],
		//			"задача №":              tasks[i%len(tasks)],
		//		}
		//		fmt.Println("Обновление кэша на", t.Format("15:04:05.000000"))
		//		stats.Update(newCache)
		//		i++
		//	}
		//}()
		//
		//// Чтение статистики из нескольких горутин
		//for i := 0; i < 100; i++ {
		//	wg.Add(1)
		//	go func(id int) {
		//		defer wg.Done()
		//		for j := 0; j < 50; j++ {
		//			lang := stats.GetStatistics("язык программирования")
		//			task := stats.GetStatistics("задача №")
		//			fmt.Printf("В %v участник №%d-%d изучает %v, задача №%v\n",
		//				time.Now().Format("15:04:05.000000"), id, j, lang, task)
		//			time.Sleep(1 * time.Second)
		//		}
		//	}(i)
		//}
		//wg.Wait()
		//fmt.Println("Все читатели завершились. Завершаем программу.")
	}
	//Задача 8.
	{
		//shutdown := &goroutine2.GracefulShutdown{}
		//// Запускаем рабочие горутины
		//for i := 0; i < 20; i++ {
		//	wg.Add(1)
		//	go func(id int) {
		//		defer wg.Done()
		//		for {
		//			if shutdown.GetShutdown() { //Бесконечный цикл
		//				fmt.Printf("Горутина %d завершает работу\n", id)
		//				return
		//			}
		//			fmt.Printf("Горутина %d работает...\n", id)
		//			time.Sleep(1 * time.Second)
		//		}
		//	}(i)
		//}
		//// Спустя 5 секунд активируем shutdown
		//time.Sleep(5 * time.Second)
		//wg.Add(1)
		//go func() {
		//	defer wg.Done()
		//	shutdown.SetShutdown()
		//}()
		//wg.Wait()
		//fmt.Println("Программа завершена")
	}
	//Задача 9. Хранилище временных меток
	{
		//tagStorage := goroutine2.NewEvents()
		//num := 15
		//for i := 0; i < 10; i++ {
		//	wg.Add(1)
		//	go func(id int) {
		//		defer wg.Done()
		//		for j := 0; j < num; j++ { //
		//			get := tagStorage.GetEvents("метка")
		//			fmt.Printf("Горутина %d-%d читает метку на период %v.\n", id, j, get.Format("15:04:05"))
		//			time.Sleep(1 * time.Second)
		//		}
		//	}(i)
		//
		//}
		//// Периодическое обновление одной метки
		//wg.Add(1)
		//go func() {
		//	defer wg.Done()
		//	for i := 0; i < num/2; i++ {
		//		tagStorage.SetEvents("метка")
		//		time.Sleep(2 * time.Second)
		//	}
		//}()
		//wg.Wait()
	}
	//Задача 10. Rate limiter на atomic
	{
		limiter := goroutine2.NewRateLimiter(9, time.Second)
		limiter.RunAutoReset()
		defer limiter.Stop()

		for i := 0; i < 20; i++ { // Количество запросов
			wg.Add(1)
			go func(id int) {
				defer wg.Done()
				if limiter.Allow() {
					fmt.Printf("Запрос %d принят\n", id)
				} else {
					fmt.Printf("Запрос %d отклонён (лимит)\n", id)
				}
				time.Sleep(300 * time.Millisecond)
			}(i)
		}
		wg.Wait()
		fmt.Println("Тест завершен")
	}
}
