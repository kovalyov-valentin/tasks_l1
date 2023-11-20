package main

import (
	"fmt"
	"math/big"
)

// Функция реализующая калькулятор
func Calculator(a, b *big.Int, operator string) *big.Int {
	sum := new(big.Int) // Переменная в которой хранится сумма
	switch operator { // Определение операторов и производим необходимые операции
	case "+":
		return sum.Add(a, b)
	case "-":
		return sum.Sub(a, b)
	case "*":
		return sum.Mul(a, b)
	case "/":
		return sum.Div(a, b)
	}
	return sum
}

func main() {
	// Создаем длинные числа
	a := new(big.Int)
	b := new(big.Int)

	// Инициализируем
	a.SetString("7218882428714186575617218882428714186575617", 10,)
	b.SetString("922337203685477580792233720368547758079223", 10,)

	// Выводим на экран числа
	fmt.Printf("a = %s\n", a.Text(10))
	fmt.Printf("b = %s\n", b.Text(10))

	// Выводим результат
	fmt.Printf("a + b = %s\n", Calculator(a, b, "+"))
	fmt.Printf("a - b = %s\n", Calculator(a, b, "-"))
	fmt.Printf("a * b = %s\n", Calculator(a, b, "*"))
	fmt.Printf("a / b = %s\n", Calculator(a, b, "/"))
}