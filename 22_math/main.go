// Разработать программу, которая перемножает, делит, складывает, вычитает
// две числовых переменных a,b, значение которых > 2^20.

// 2^20 — это 1 048 576
// Для работы с числами произвольной длины,
// не ограниченной размерами фиксированных типов данных,
// воспользуемся стандартной библиотекой math/big

package main

import (
	"bufio"
	"fmt"
	"math/big"
	"os"
	"strings"
)

func main() {

	// Получаем операцию
	fmt.Println("Please enter your expression below:") // например: 1148576 / 1258376
	first, operation, second := getExpression()

	// Создаем большие числа a и b
	a, ok := new(big.Float).SetString(first)
	if !ok {
		fmt.Println("Failed to set 'a'")
		return
	}
	b, ok := new(big.Float).SetString(second)
	if !ok {
		fmt.Println("Failed to set 'b'")
		return
	}

	operations := "+-*/"

	contains := strings.Contains(operations, operation)

	if !contains {
		fmt.Println("Wrong operation")
		return
	}

	result := new(big.Float)

	switch operation {
	case "*":
		// Умножение
		result.Mul(a, b)
		fmt.Println("Умножение:", result)
	case "/":
		// Деление
		result.Quo(a, b)
		fmt.Println("Деление:", result)
	case "+":
		// Сложение
		result.Add(a, b)
		fmt.Println("Сложение:", result)
	case "-":
		// Вычитание
		result.Sub(a, b)
		fmt.Println("Вычитание:", result)
	}
}

func getExpression() (a, operation, b string) {

	// Создаём сканнер
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()
	str := scanner.Text()

	// Валидируем ввод
	if str == "" {
		fmt.Println("String can't be empty")
		return getExpression()
	}
	words := strings.Split(str, " ")
	if len(words) != 3 {
		fmt.Println("Not enough/too much arguments")
		fmt.Println("Tip: operands and operation must be separated with spaces")
		return getExpression()
	}

	a = words[0]
	operation = words[1]
	b = words[2]
	return a, operation, b
}
