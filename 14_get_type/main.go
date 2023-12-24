// Разработать программу, которая в рантайме способна определить тип переменной:
// int, string, bool, channel из переменной типа interface{}.
package main

import (
	"fmt"
)

func getType(value interface{}) {
	switch v := value.(type) {
	case int:
		fmt.Println("Тип int, значение", v)
	case string:
		fmt.Println("Тип string, значение", v)
	case bool:
		fmt.Println("Тип bool, значение", v)
	case chan int:
		fmt.Println("Тип chan int")
	default:
		fmt.Printf("Unexpected type %T\n", v)
	}
}

func main() {

	getType(42)
	getType("hello")
	getType(true)

	channel := make(chan int)
	getType(channel)

	// Если полученный тип не ожидался
	getType(14.7)
}
