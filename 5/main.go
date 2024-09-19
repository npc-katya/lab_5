package main

import (
	"fmt"
	"math"
)

// интерфейс Shape
type Shape interface {
	Area() float64
}

// структура Circle
type Circle struct {
	radius int
}

// метод структуры Circle
func (circle Circle) Area() (area float64) {
	var radius float64
	radius = float64(circle.radius)
	area = radius * radius * math.Pi
	return area
}

// структура Rectangle
type Rectangle struct {
	x int
	y int
}

// метод структуры Rectangle
func (rectangle Rectangle) Area() float64 {
	area := rectangle.x * rectangle.y
	return float64(area)
}

// функция для вывода площади каждого объекта Shape
func printAreas(shapes []Shape) {
	for _, shape := range shapes {
		fmt.Printf("площадь: %.2f\n", shape.Area())
	}
}

func main() {

	// Реализовать функцию, которая принимает срез интерфейсов Shape и выводит площадь каждого объекта.

	// создание среза объектов Shape
	shapes := []Shape{
		Circle{radius: 5},
		Rectangle{x: 4, y: 6},
	}

	// вывод площади
	printAreas(shapes)

}
