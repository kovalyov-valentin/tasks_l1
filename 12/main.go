package main

import "fmt"

// Проверка на наличие значения в слайсе
func contains(strArray []string, str string) bool{
	if len(strArray) == 0 { // Проверка пустой слайс или нет
		return true
	}

	for _, s := range strArray {
		if s == str { // Если стринга уже есть, то мы ее не добавляем
			return false
		}
	}

	return true
}

// Создаем собственное множество для нашей последовательности строк
func Set(str []string) {
	var result []string // Здесь у нас конечное значение
	for _, s := range str {
		if contains(result, s) { // Добавляем значение, если оно не содержится в конечном слайсе
			result = append(result, s)
		}
	}
	fmt.Println(result)
}

func main() {
	str := []string{"cat", "cat", "dog", "cat", "tree"}
	Set(str)
}