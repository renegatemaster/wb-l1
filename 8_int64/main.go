// Дана переменная int64. Разработать программу,
// которая устанавливает i-й бит в 1 или 0.
package main

import (
	"fmt"
)

// Устанавливает значение бита в числе
func setBit(data int64, bit int, value int) int64 {
	if value == 1 {
		// Устанавливает i-й бит в 1
		return data | (1 << bit)
	} else {
		// Устанавливает i-й бит в 0
		return data &^ (1 << bit)
	}
}

// Получаем из ввода данные
func getData() (int64, int, int) {

	var (
		data  int64
		bit   int
		value int
	)

	fmt.Println("Enter int64:")
	fmt.Scanln(&data)

	fmt.Println("Enter the bit you want to change:")
	fmt.Scanln(&bit)
	if bit < 0 || bit > 63 {
		fmt.Println("Bit out of range\nPlease enter valid value:")
		return getData()
	}

	fmt.Println("Enter the value (1 or 0):")
	fmt.Scanln(&value)
	if value < 0 || value > 1 {
		fmt.Println("Value out of range\nPlease enter 1 or 0:")
		return getData()
	}

	return data, bit, value
}

func main() {

	data, bit, value := getData()

	data = setBit(data, bit, value)
	fmt.Printf("After setting the %d bit to %d: %d\n", bit, value, data)
}
