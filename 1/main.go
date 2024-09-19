package main

import "fmt"

// структура с полями name, age
type Person struct {
	name string
	age  int
}

// метод, который выводит информацию
func (person Person) info() {
	fmt.Println("имя:", person.name)
	fmt.Println("возраст:", person.age)
}

func main() {

	// Создать структуру Person с полями name и age. Реализовать метод для вывода информации о человеке.

	// создание экземпляра
	alice := Person{name: "Алиса", age: 25}

	// вызов метода
	alice.info()

}
