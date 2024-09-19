package main

import (
	"fmt"
)

// интерфейс Stringer
type Stringer interface {
	String() string
}

// структура Book
type Book struct {
	Title  string
	Author string
}

// метод структуры Book
func (b Book) String() string {
	return fmt.Sprintf("книга: %s - %s", b.Title, b.Author)
}

// метод интерфейса Stringer
func Info(s Stringer) {
	fmt.Println(s.String())
}

func main() {

	// Создать интерфейс Stringer и реализовать его для структуры Book, которая хранит информацию о книге.

	// инициализация объекта Book и передача в Info()
	book := Book{"Мёртвые души", "Николай Гоголь"}
	Info(book)
}
