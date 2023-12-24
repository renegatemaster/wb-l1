package main

import (
	"fmt"
	"os"
	"sync"
)

func main() {

	var array [5]int = [5]int{2, 4, 6, 8, 10}

	// Конкурентный расчёт предполагает использование горутин
	// Главная горутина — функция main()
	// Чтобы она не отработала до исполнения всех горутин,
	// используем WaitGroup для синхронизации
	wg := sync.WaitGroup{}

	for _, num := range array {
		wg.Add(1) // добавляем ожидание каждой горутины
		go func(num int) {
			fmt.Fprintln(os.Stdout, num*num)
			wg.Done() // сигнал о завершении горутины
		}(num)
	}

	wg.Wait() // Ждём завершения горутин
}
