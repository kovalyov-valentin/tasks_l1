package main

import (
	"fmt"
	"sort"
)



func binarySearch(nums []int, elem int) int {
	l := len(nums)

	low := 0
	high := l - 1

	if l < 2 { // Если массив состоит из одного элемента то возвращаем elem
		return elem
	}

	for low <= high {
		middle := (low + high) / 2 // Середина слайса                          
		
		guess := nums[middle]

		if guess == elem {
			return middle
		}

		if guess < elem { // Если элемент, который мы ищем меньше, то ищем в левой половине, если больше то в правой
			low = middle + 1
		} else {
			high = middle - 1
		}
	}

	return -1
}

func main() {
	nums := []int{95,78,46,58,45,86,99,251,320}
	sort.Ints(nums) // Для бинарного поиска наш массив должен быть отсортирован
	fmt.Println("Sort arrays:", nums)
	fmt.Println("Index value:",binarySearch(nums, 58))
}