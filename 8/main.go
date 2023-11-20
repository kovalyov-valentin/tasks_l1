package main

import "fmt"

// SetBit устанавливает или сбрасывает бит в числе num по индексу i
func SetBit(num int64, i int, set bool) int64 {
	if set {
		// Устанавливаем бит, используя побитовое или
		num = num | (1 << i)
	} else {
		// Сбрасываем бит, используя побитовое исключающее или
		num = num ^ (1 << i)
	}
	return num
}

func main() {
	var num int64 
	// Устанавливаем 10-й бит в единицу
	num = SetBit(num, 10, true)
	fmt.Println(num)

	// Устанавливаем 2-й бит в единицу
	num = SetBit(num, 2, true)
	fmt.Println(num)

	// Сбрасываем 10-й бит в ноль
	num = SetBit(num, 10, false)
	fmt.Println(num)

	// Устанавливаем 0-й бит в единицу.
	num = SetBit(num, 0, true)
	fmt.Println(num)
}