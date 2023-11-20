package main

import "fmt"

// Функция для определения типа
func Type(i interface{}) { 
	switch v := i.(type) { //Ключевое слово type помогает нам получить тип элемента
	case int:
		fmt.Printf("Type: int, value: %d\n", v) // И просто выводим тип нашего значения
	case float64:
		fmt.Printf("Type: float, value: %v\n", v)
	case string:
		fmt.Printf("Type: string, value: %v\n", v)
	case bool:
		fmt.Printf("Type: bool, value: %v\n", v)
	case chan int:
		fmt.Printf("Type: chan, value: %d\n", v)
	default:
		fmt.Printf("Undefined type for value: %v\n", v)
	}
}

func main() {
	elem := []interface{}{1, 15.4, "hello", true, make(chan int)}

	for _, e := range elem { // Проходимся по нашему слайсу и передаем значения в функцию для определения типа
		Type(e)
	}
}
