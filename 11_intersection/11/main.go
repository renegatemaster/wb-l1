// Реализовать пересечение двух неупорядоченных множеств.

// В Golang нет встроенного типа данных для неупорядоченных множеств,
// Используем map для эмуляции неупорядоченного множества,
// где ключи служат элементами множества, а значения не имеют значения.
package main

import (
	"fmt"
)

func main() {

	first := map[int]bool{
		1:  true,
		0:  true,
		17: true,
		28: true,
		84: true,
		42: true,
		78: true,
	}

	second := map[int]bool{
		4:  true,
		78: true,
		10: true,
		42: true,
		66: true,
		17: true,
	}

	// Получаем новый неупорядоченный массив из пересечений
	third := intersect(first, second)

	// Выводим значения пересечений в одну строку
	for i := range third {
		fmt.Printf("%d ", i)
	}
	fmt.Println()
}

func intersect(first, second map[int]bool) map[int]bool {

	third := map[int]bool{}

	for i := range first {
		for j := range second {
			switch {
			case i == j:
				third[i] = true
			default:
				continue
			}
		}
	}

	return third
}
