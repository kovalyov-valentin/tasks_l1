package main

import "fmt"

func MethodSwap(a, b int) {
	a, b = b, a 
	fmt.Println("Method Swap:", a, b)
}


// Стоит учесть, что XOR не работает c float 
func MethodXOR(a, b int) {
	a = a ^ b
	b = a ^ b 
	a = a ^ b
	fmt.Println("Method XOR:", a, b)
}

func MethodPlusMinus(a, b int) {
	a = a + b 
	b = a - b 
	a = a - b 
	fmt.Println("Method PlusMinus", a, b)
}

func main() {
	a := 10
	b := 20 
	fmt.Println("Original value:", a, b)
	MethodSwap(a, b)
	MethodXOR(a, b)
	MethodPlusMinus(a, b)
}