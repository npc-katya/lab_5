package main

import (
	"fmt"
	"math"
)

// структура Circle с полем radius
type Circle struct {
	radius int
}

// функция, вычисляет площадь круга
func (circle Circle) Area() (area float64) {
	var radius float64
	radius = float64(circle.radius)
	area = radius * radius * math.Pi
	return area
}

func main() {

	// Создать структуру Circle с полем radius и метод для вычисления площади круга.

	// создание экземпляра
	сircle := Circle{radius: 13}

	// вызов метода
	fmt.Println("радиус круга:", сircle.radius)
	fmt.Println("площадь круга:", сircle.Area())

}
