package main

import (
	"fmt"
	"sync"
)

// Рассчитываем значение квадратов чисел
func square(num int) int {
	return num * num
}

// Решение при помощи Channel
func ConcurrentSquareChannel(nums []int) {

	var result int

	ch := make(chan int, len(nums)) // Инициализируем канал
	for _, n := range nums {
		go func(n int) {
			ch <- square(n) // Запись значений в канал
		}(n)
	}

	for i := 0; i < len(nums); i++ { // Считывание значений из канала
		result = <-ch // Записываем в результат
		fmt.Println(result)
	}
	close(ch) // Закрываем канал, т.к через него больше не будут отправляться значения

}

// Решение при помощи WairGroup и Mutex
func ConcurrentSquareWG(nums []int) {
	wg := sync.WaitGroup{}
	mu := sync.Mutex{}

	for _, n := range nums {
		wg.Add(1)        // Добавляем элементы в очередь
		go func(n int) { // Здесь мы прокидываем значеие в функцию
			defer wg.Done()
			mu.Lock() // Происходит блокировка доступа к ресурсу
			fmt.Println(square(n))
			mu.Unlock()
		}(n)
	}

	wg.Wait()
}

func main() {
	nums := []int{2, 4, 6, 8, 10}

	fmt.Println("Starting ConcurrentSquareChannel:")
	ConcurrentSquareChannel(nums)
	fmt.Println("Starting ConcurrentSquareWG:")
	ConcurrentSquareWG(nums)
}
