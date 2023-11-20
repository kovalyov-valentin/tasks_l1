package main

import "fmt"

func Reverse(str string) string {
	s := []rune(str) // Преобразуем строку в слайс рун, так как символы могут быть unicode

	l := len(s) // Просто определяем длину нашей стринги

	// Индекты для первого и последнего символа строки
	i := 0
	j := l - 1

	for i < j { // Проходимся по строке с двух сторон
		s[i], s[j] = s[j], s[i] // Свапаем символы местами
		i++
		j--
	}

	return string(s)
}

func main() {
	str := "главрыба"
	s := Reverse(str)
	fmt.Println(s)
}
