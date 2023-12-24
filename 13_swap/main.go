// Поменять местами два числа без создания временной переменной.
package main

import "fmt"

func main() {

	first := 1
	second := 2

	fmt.Printf("Переменная first = %d, переменная second = %d\n", first, second)

	first, second = second, first

	fmt.Printf("Переменная first = %d, переменная second = %d\n", first, second)
}
