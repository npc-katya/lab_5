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

// метод интерфейса Shape
func Area(s Shape) {
	fmt.Println("площадь:", s.Area())
}

func main() {

	// Создать интерфейс Shape с методом Area(). Реализовать этот интерфейс для структур Rectangle и Circle.

	// инициализация объекта Circle и передача в Area()
	circle := Circle{radius: 3}
	fmt.Println("круг с радиусом", circle.radius)
	Area(circle)

	fmt.Println()

	// инициализация объекта Rectangle и передача в Area()
	rectangle := Rectangle{x: 3, y: 4}
	fmt.Println("прямоугольник со сторонами", rectangle.x, rectangle.y)
	Area(rectangle)

}
