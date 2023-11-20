package main

import (
	"fmt"
	"math"
)

// Структура Point
type Point struct {
	x float64
	y float64
}

// Конструктор для создания нового экземпляра Point
func NewPoint(x, y float64) *Point {
	return &Point{
		x: x,
		y: y,
	}
}

// Функция, использующая формулу для нахождения расстояния между точками
func Distance(a, b *Point) float64 {
	return math.Sqrt(math.Pow((b.x-a.x), 2) + math.Pow((b.y - a.y), 2))
}

func main() {
	a := NewPoint(4.3, 7.6)
	b := NewPoint(8.1, 3.4)
	fmt.Println(Distance(a, b))
}
