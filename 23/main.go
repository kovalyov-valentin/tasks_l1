package main

import "fmt"

// Удаление элемента по индексу из слайса
func DeleteByIndex(nums []int, index int) []int {
	if index < 0 || index >= len(nums) { // Проверяем индекс на корректность
		return nil
	}

	if index != len(nums)-1 { // Проверяем не является ли элемент, который мы хотим удалить, последним
		return append(nums[:index], nums[index+1:]...) // Создаем новый слайс, состоящий из всех элементов до удаляемого элемента и всех элементов после
	} else {
		return nums[:index]
	}
}

func main() {
	nums := []int{2, 4, 6, 8, 1, 3, 5, 7, 9}
	fmt.Println("Original slice:", nums)
	delete := DeleteByIndex(nums, 5)
	fmt.Println("Slice after delete:", delete)
}
