package main

import (
	"fmt"
	"os"
)

func main() {
	var array [5]int = [5]int{2, 4, 6, 8, 10}

	// Используем буферизированный канал для синхронизации горутин
	ch := make(chan int, len(array))

	for _, num := range array {
		go func(num int) {
			ch <- (num * num) // результат записывается в канал
		}(num)
	}

	for i := 0; i < len(array); i++ {
		fmt.Fprintln(os.Stdout, <-ch)
	}
}
