package main

import (
	"fmt"
	"os"
)

func main() {
	var array [5]int = [5]int{2, 4, 6, 8, 10}

	// Создаём буфферизованный канал длиной в массив
	ch := make(chan int, len(array))

	for _, num := range array {
		go func(num int) {
			ch <- num * num
		}(num)
	}

	// Значение не выводится пока не будут прочитаны необходимые данные из канала
	fmt.Fprintln(os.Stdout, (<-ch + <-ch + <-ch + <-ch + <-ch))
}
