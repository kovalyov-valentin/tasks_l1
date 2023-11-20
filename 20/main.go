package main

import (
	"fmt"
	"strings"
)

func ReverceWord(str string) string {
	words := strings.Fields(str) // Разбиваем строку и возвращаем фрагмент подстрок

	l := len(words) // Просто определяем длину 

	// Индекты для первого и последнего символа строки
	i := 0
	j := l -1 

	for  i < j{ // Проходимся по строке с двух сторон
		words[i], words[j] = words[j], words[i] // Свапаем символы местами
		i++
		j--
	}

	return strings.Join(words, " ") // Создаем единую строку, разделяя слова пробелом
}

func main() {
	words := "snow dog sun"
	fmt.Println("Original string:", words)
	reverce := ReverceWord(words)
	fmt.Println("Reverce string:", reverce)


}