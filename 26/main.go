package main

import (
	"fmt"
	"strings"
)

// Проверка строки на уникальность
func UniqueSymbols(str string) bool {
	str = strings.ToLower(str) // Перевод символов в нижний регистр

	m := make(map[rune]int) // Будем помещать значения в мапу

	for _, s := range str { // Проходимся по символам строки
		if _, ok := m[s]; !ok { // Все ок если ключ с текущим значением не найден
			m[s] = 0
		} else {
			return false // Если ключ с текущим значением найден, значит символ уже встречался
		}
	}
	return true
}

func main() {
	str := []string{"abcd", "abCdefAaf", "aabcd"}

	for _, s := range str {
		if UniqueSymbols(s) {
			fmt.Printf("%s unique\n", s)
		} else {
			fmt.Printf("%s not unique\n", s)

		}
	}
}