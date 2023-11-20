package main

import "fmt"

// Структура Human
type Human struct {
	FirstName string
	LastName  string
	Age       int
}

// Структура Action
type Action struct {
	Human // Встроенная структура Human, теперь мы можем использовать методы Human
}

func (h Human) Info() { // Метод структуры Human
	fmt.Printf("Hello! My name is %s %s\n", h.FirstName, h.LastName)
}

func (a Action) Programming() { // Метод структуры Action
	fmt.Printf("I seriously started programming at the age of %d\n", a.Age)
}


func main() {
	person := Action{
		Human{
			FirstName: "Valentin",
			LastName: "Kovalyov",
			Age: 24,
		},
	}

	person.Info() // Использование метода встроенной структуры Human
	person.Programming() // Вызов собственного метода Action
}
