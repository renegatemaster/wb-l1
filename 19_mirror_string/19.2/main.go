// Разработать программу, которая переворачивает подаваемую на ход строку
// (например: «главрыба — абырвалг»).
// Символы могут быть unicode.

package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {

	// Получаем строку
	initial := getString()

	mirrored := mirrorString(initial)

	fmt.Printf("«%s — %s»\n", initial, mirrored)
}

func mirrorString(input string) string {

	// Преобразуем строку в слайс рун
	runes := []rune(input)

	// Переворачиваем слайс
	for i, j := 0, len(runes)-1; i < j; i, j = i+1, j-1 {
		runes[i], runes[j] = runes[j], runes[i]
	}

	// Формируем строку из слайса
	return string(runes)
}

func getString() string {
	fmt.Println("Enter string:")
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()
	str := scanner.Text()
	if str == "" {
		fmt.Println("String can't be empty")
		return getString()
	}
	return str
}
