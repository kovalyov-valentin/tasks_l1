// Паттерн "Адаптер", который я ранее реализовывал по книге "Head First: Паттерны проектирования".

package main

import (
	"fmt"
)

// duck интерфейс определяет методы quack() и fly()
type duck interface {
	quack()
	fly()
}

// mallarDuck структура представляет объект утки и реализует интерфейс duck.
type mallarDuck struct{}

func (m *mallarDuck) quack() {
	fmt.Println("Quack")
}

func (m *mallarDuck) fly() {
	fmt.Println("I'm flying")
}

// turkey интерфейс определяет методы gobble() и fly().
type turkey interface {
	gobble()
	fly()
}

// wildTurkey структура представляет объект индейки и реализует интерфейс turkey.
type wildTurkey struct{}

func (w *wildTurkey) gobble() {
	fmt.Println("Gobble Gobble")
}

func (w *wildTurkey) fly() {
	fmt.Println("I'm flying a short distance")
}

// turkeyAdapter структура представляет адаптер для индейки, реализующий интерфейс duck.
type turkeyAdapter struct {
	turkey turkey
}

// newTurkeyAdapter функция создает новый объект turkeyAdapter.
func newTurkeyAdapter(t turkey) *turkeyAdapter {
	return &turkeyAdapter{
		turkey: t,
	}
}

func (t *turkeyAdapter) quack() {
	t.turkey.gobble()
}

func (t *turkeyAdapter) fly() {
	for i := 1; i <= 2; i++ {
		t.turkey.fly()
	}
}

func main() {
	// Создание утки и индейки
	mallarDuck := &mallarDuck{}
	wildTurkey := &wildTurkey{}

	// Заворачиваем индейку в адаптер для индейки, что делает ее похожей на утку
	turkeyAdapter := newTurkeyAdapter(wildTurkey)

	//Передача duck методу testDuck(), который ожидает объект Duck
	fmt.Println("The Duck says...")
	testDuck(mallarDuck)

	// Выдать индейку за утку тем же способом
	fmt.Println("\nThe Turkey Adapter says...")
	testDuck(turkeyAdapter)
}

// testDuck функция принимает объект duck и вызывает его методы quack() и fly().
func testDuck(d duck) {
	d.quack()
	d.fly()
}
