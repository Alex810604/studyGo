package main

import (
	"awesomeProject/goroutines/task_goroutunes/task"
	"fmt"
	"sync"
	"sync/atomic"
)

var wg sync.WaitGroup
var mu sync.Mutex

func main() {
	//Задача 1.
	//go goroutine.SimpleGoroutine()
	//time.Sleep(time.Millisecond)

	////Задача 2.
	//for i := 1; i <= 5; i++ {
	//	wg.Add(1)
	//	go goroutine.LotsGoroutines(i, &wg)
	//}
	//wg.Wait()
	//fmt.Println("Все горутины выполнены")

	//Задача 3.
	//wg.Add(1)
	//go goroutine.GoroutineDelay(&wg)
	//fmt.Println("Ждем завершения горутины...")
	//wg.Wait()

	//Задача 4.
	//ch := make(chan int)
	//func() {
	//	for v := range ch {
	//		fmt.Println("Чтение из канала", v)
	//	}
	//}()
	//for i := 0; i < 3; i++ {
	//	wg.Add(2)
	//	go func(id int) {
	//		defer wg.Done()
	//		ch <- id
	//		fmt.Println("канал", id)
	//	}(i)
	//	go goroutine.CycleGoroutine(i, &wg)
	//}
	//wg.Wait()
	//close(ch)

	//Задача 5.
	//var counter int
	//for i := 0; i < 1000; i++ {
	//	wg.Add(1)
	//	go goroutine.DataRace(&wg, i, &counter)
	//}
	//wg.Wait()
	//fmt.Println("Итоговый счетчик", counter)

	//Задача 6.
	//var counter1 int
	//for i := 0; i < 1000; i++ {
	//	wg.Add(1)
	//	go goroutine.DataRaceMutex(&wg, &mu, i, &counter1)
	//}
	//wg.Wait()
	//fmt.Println("Итоговый счетчик", counter1)

	//Задача 7.
	//var counter2 int
	//for i := 1; i <= 4; i++ {
	//	wg.Add(1)
	//	go goroutine.DataWait(&wg, i, &counter2)
	//}
	//wg.Wait()
	//fmt.Println("Итоговый счетчик", counter2)

	//Задача 8.
	//var result int64 // Разделяемая переменная
	//n := 5
	//wg.Add(1)
	//go goroutine.NumberSquared(&wg, &result, n)
	//wg.Wait()
	//fmt.Printf("Квадрат числа %d равен %d.\n", n, atomic.LoadInt64(&result))
	//fmt.Printf("Квадрат числа %d равен %d.\n", n, result)

	//Задача 9.
	var result int32 // Разделяемая переменная
	num := 5
	wg.Add(2)
	go func() {
		defer func() {
			if r := recover(); r != nil {
				fmt.Println("Паника перехвачена:", r)
			}
			wg.Done()
		}()
		goroutine.NumberCubed(&wg, num, &result)
		fmt.Println("В работе...")
	}()
	wg.Wait()
	fmt.Printf("Куб числа %d равен %d.", num, atomic.LoadInt32(&result))

	//Задача 10.
	//var flag atomic.Bool
	//wg.Add(1)
	//go goroutine.InfiniteLoop(&wg, &flag)
	//time.Sleep(3 * time.Second)
	//flag.Store(true)
	//wg.Wait()
}
