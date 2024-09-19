package main

import "fmt"

// структура с полями name, age
type Person struct {
	name string
	age  int
}

// метод, который увеличивает возраст на 1 год
func (person *Person) Birthday() {
	person.age++
}

func main() {

	// Реализовать метод birthday для структуры Person, который увеличивает возраст на 1 год.

	// создание экземпляра
	alice := Person{name: "Алиса", age: 25}

	fmt.Println("возраст до birthday:", alice.age)

	// вызов метода
	alice.Birthday()
	fmt.Println("возраст после birthday:", alice.age)

}
