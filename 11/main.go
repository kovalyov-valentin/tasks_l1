package main

import "fmt"

// Функция, которая проверяет присутствует ли значение в слайсе
func contains(nums []int, num int) bool { 
	for _, n := range nums {
		if n == num { // Если в слайсе присутствует значение, то у нас true и двигаемся дальше
			return true
		}
	}

	return false
}

func Intersection(firstNum, secondNum []int) {
	var result []int // Слайс, который будем содержать пересечения других слайсов
	for _, v := range firstNum { // Проходимся по первому слайсу
		if contains(secondNum, v) { // Если значение присутствует во втором слайсе, то добавляем его в результирующий слайс
			result = append(result, v)
		}
	}
	fmt.Println(result)
}

func main() {
	firstNum := []int{1, 3, 5, 7, 9}
	secondNum := []int{2, 4, 6, 8, 5}
	Intersection(firstNum, secondNum)
}
