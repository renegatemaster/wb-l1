// Разработать программу, которая переворачивает подаваемую на ход строку
// (например: «главрыба — абырвалг»).
// Символы могут быть unicode.

package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {

	// Получаем строку
	initial := getString()

	mirrored := mirrorString(initial)

	fmt.Printf("«%s — %s»\n", initial, mirrored)
}

func mirrorString(str string) string {

	// Создаём слайс символов строки
	signs := strings.Split(str, "")
	// Создаём пустую строку для записи результата
	var result string

	// Записываем в пустую строку символы в обратном порядке
	for i := len(signs) - 1; i >= 0; i-- {
		result += signs[i]
	}

	return result
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
