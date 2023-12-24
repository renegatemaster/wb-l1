/*
Разработать программу нахождения расстояния между двумя точками,
которые представлены в виде структуры Point
с инкапсулированными параметрами x,y и конструктором.

Расстояние между двумя точками на плоскости равно
квадратному корню из суммы квадратов разностей
соответствующих координат:

d = √( (x2-x1)^2 +(y2-y1)^2 )
*/

package main

import (
	"fmt"
	"math"
)

type Point struct {
	x, y float64
}

func NewPoint(x, y float64) *Point {
	return &Point{x, y}
}

func distance(first, second *Point) float64 {

	// Рассчитываем квадрат разности по осям
	xDimension := math.Pow(second.x-first.x, 2)
	yDimension := math.Pow(second.y-first.y, 2)

	// Возвращаем квадратный корень суммы квадратов разностей соответсвующих осей
	return math.Sqrt(xDimension + yDimension)
}

func main() {
	first := NewPoint(1, 1)
	second := NewPoint(9, 9)
	fmt.Println(distance(first, second))
}
