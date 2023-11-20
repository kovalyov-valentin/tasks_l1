package main

import "fmt"

// NumsFromArrays принимает срез чисел и возвращает канал для отправки этих чисел
func NumsFromArrays(nums []int) <-chan int {

	// Создаем канал для отправки чисел
	out := make(chan int)

	// Запускаем горутину для отправки чисел в канал
	go func() {
		defer close(out)

		// Отправляем каждое число из среза в канал
		for _, n := range nums {
			out <- n
		}
	}()

	return out
}

// Square принимает канал с числами, возводит каждое число в квадрат и возвращает новый канал
func Square(in <- chan int) <- chan int {
	out := make(chan int)
	go func() {
		defer close(out)
		// Читаем числа из входного канала и отправляем их квадраты в новый канал
		for n := range in {
			out <- n * n
		}
	}()

	return out
}

// Print принимает канал с числами и выводит их на экран
func Print(in <- chan int) {
	for num := range in {
		// Читаем числа из канала и выводим их на экран
		fmt.Println(num)
	}
}

func main() {
	nums := []int{2, 4, 6, 8, 1, 3, 5, 7, 9}

	// Создаем канал для отправки чисел из среза
	num := NumsFromArrays(nums)
	// Создаем канал для отправки квадратов чисел
	out := Square(num)
	// Выводим квадраты чисел
	Print(out)
	
}