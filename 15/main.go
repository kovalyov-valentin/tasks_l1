package main

import (
	"fmt"
	"strings"
	"unicode/utf8"
)

var justString string

func someFunc() {
	v := createHugeString(1 << 10)
	// Руна может занимать больше одного байта
	fmt.Println(utf8.RuneLen('п')) // 2 байта. Могло быть и больше, от 1 до 4 байт

	justString = v[:100] // В данном примере происходит срез по количеству байт, а не по рунам

	r := []rune(v) // Преобразуем стрингу в срез рун

	justString = string(r[:100])
}

func createHugeString(lenght int) string {
	var str strings.Builder // Используем для эффективной конкатенации строк

	for i := 0; i < lenght; i++ {
		fmt.Fprintf(&str, "п")
	}

	return str.String()
}

func main() {
	someFunc()
}

// Фрагмент кода на языке Go, приведенный в тексте задания, может привести к утечке памяти из-за использования
// среза v[:100] для присвоения значения переменной justString. При этом ссылка на огромную строку в памяти
// сохраняется, и даже если она больше не используется, сборщик мусора не освободит её,
// так как ссылка на огромную строку сохранена в justString.
