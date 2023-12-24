// Разработать конвейер чисел.
// Даны два канала: в первый пишутся числа (x) из массива,
// во второй — результат операции x*2,
// после чего данные из второго канала должны выводиться в stdout.
package main

import (
	"fmt"
	"os"
	"sync"
)

func main() {

	var array [5]int = [5]int{2, 4, 6, 8, 10}
	chanArray := make(chan int)
	chanSquares := make(chan int)

	wg := sync.WaitGroup{}

	wg.Add(1)
	// Записываем квадрат значений первого канала во второй
	go func() {
		for num := range chanArray {
			chanSquares <- num * 2
		}
		close(chanSquares)
		wg.Done()
	}()

	wg.Add(1)
	// Выводим на печать
	go func() {
		for num := range chanSquares {
			fmt.Fprintln(os.Stdout, num)
		}
		wg.Done()
	}()

	// Записываем значения х в первый канал
	for _, num := range array {
		chanArray <- num
	}
	close(chanArray)
	wg.Wait()
}
