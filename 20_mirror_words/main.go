// Разработать программу, которая переворачивает слова в строке.
// Пример: «snow dog sun — sun dog snow».

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

	mirrored := mirrorWords(initial)

	fmt.Printf("«%s — %s»\n", initial, mirrored)
}

func mirrorWords(str string) string {

	// Создаём слайс символов строки
	words := strings.Split(str, " ")
	var result []string

	for i := len(words) - 1; i >= 0; i-- {
		result = append(result, words[i])
	}

	return strings.Join(result, " ")
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
