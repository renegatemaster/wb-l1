/*
Разработать программу, которая проверяет,
что все символы в строке уникальные (true — если уникальные, false etc).
Функция проверки должна быть регистронезависимой.

Например:
abcd — true
abCdefAaf — false
aabcd — false
*/

package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {

	// Получаем строку
	str := getString()

	// Проверяем на уникальность
	fmt.Println(unique(str))
}

func unique(str string) bool {
	// Для регистранезависимости переводим все символы к нижнему регистру
	lowerStr := strings.ToLower(str)
	// Получаем слайс символов
	signs := strings.Split(lowerStr, "")
	// Создаём пустой слайс, куда будем записывать символы для проверки уникальности
	unique := []string{}

	for i := range signs {
		for j := range unique {
			if signs[i] == unique[j] {
				return false
			}
		}
		unique = append(unique, signs[i])
	}

	return true
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
