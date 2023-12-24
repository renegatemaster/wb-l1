// К каким негативным последствиям может привести данный фрагмент кода,
// и как это исправить? Приведите корректный пример реализации.

// Не понятно зачем делать переменную justString глобальной,
// если она используется только в одной функции.
// Это увеличивает шанс случайно записать в неё что-то
// в будущем из новой функции и увеличивает длину программы.
// var justString string
//
//	func someFunc() {
//		// Длинный и не особо читаемый способ передать int в функцию
//	    v := createHugeString(1 << 10)
//	    // Тут запись int'а уже другая и нет уверенности,
//	    // что заданный индекс находится в пределах длины строки
//	    // Если index out of range, то случится паника
//	    justString = v[:100]
//	}
//
// Не понятно зачем выносить функционал программы в отдельную функцию,
// если это единственная функция в программе, код читается хуже.
//
//	func main() {
//		 someFunc()
//	}
package main

import (
	"fmt"
)

func createHugeString(length int) string {

	var hugeString string
	for i := 0; i < length; i++ {
		hugeString += "1"
	}

	return hugeString
}

func main() {

	v := createHugeString(1024)
	fmt.Println(v)

	fmt.Println()

	// Желаемая длина обычной строки
	length := 100
	var justString string

	switch {
	case len(v) > length:
		justString = v[:length]
	case len(v) <= length:
		fmt.Printf("Строка длиной %d не может передать %d символов\n", len(v), length)
		fmt.Println("Возвращаем строку максимально возможной длины")
		justString = v
	}

	fmt.Println(justString)
}
