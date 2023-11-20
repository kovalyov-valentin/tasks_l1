package main

import "fmt"

// GroupTemperature группирует температуры в слайсе по диапазонам 10 градусов.
// Функция принимает слайс float64 и выводит результат группировки в виде мапы
func GroupTemperature(temperature []float64) {
	// Инициализируем карту для хранения группированных температур
	group := make(map[int][]float64)

	// Проходим по всем элементам слайса температур
	for _, temp := range temperature {
		// Вычисляем ключ группы, округляя температуру до десятков
		key := int(temp/10) * 10

		// Добавляем температуру в соответствующую группу
		group[key] = append(group[key], temp)
	}

	fmt.Println(group)
}

func main() {
	temperature := []float64{-25.4, -27.0, 13.0, 19.0, 15.5, 24.5, -21.0, 32.5}
	GroupTemperature(temperature)
}