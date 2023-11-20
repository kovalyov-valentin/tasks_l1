package main

import (
	"fmt"

)

func QuickSort(input []int) []int {
	l := len(input)

	if l < 2 { // Проверяем длину массива. Если его длина меньше 2, то считаем его отсортированным и возращаем без изменений
		return input
	}


	less := make([]int, 0) // Этот массив будет содержать элементы, которые меньше pivot
	bigger := make([]int, 0) // Этот массив будет содержать элементы, которые больше pivot
	pivot := input[0] // Первый элемент массива, который мы выбираем в качестве центрального

	for _, v := range input[1:] { // Начинаем проход по массиву по первому элементу после опорного
		if v > pivot {
			bigger = append(bigger, v) // Если элемент больше центрального
		} else {
			less = append(less, v) // Если элемент меньше центрального
		}
	}
	// Используем рекурсию нашей функции, для массивов less, bigger 
	input = append(QuickSort(less), pivot) // Результат сортировки less добавляем к pivot 
	input = append(input, QuickSort(bigger)...) // Затем к pivot добавляем результат сортировки bigger 

	return input
}

func main() {
	nums := []int{5, 4, 1, 2, 0, 123, 1234, 32, 12, 2345, 99}
	fmt.Println("Original arrays:", nums)
	
	fmt.Println("Sort using QuickSort:", QuickSort(nums))
	
	// sort.Ints(nums)
	// fmt.Println("Sort using package sort:", nums)

	
}