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
func ConcurrentSquareSumChannel(nums []int) {

	var result int
	ch := make(chan int, len(nums)) // Инициализируем канал

	for _, n := range nums {
		go func(n int) {
			ch <- square(n) // Запись значений в канал
		}(n)
	}

	for i := 0; i < len(nums); i++ { // Считывание значений из канала
		result += <-ch // Суммируем значения

	}
	close(ch) // Закрываем канал, т.к через него больше не будут отправляться значения
	fmt.Println(result)

}

// Решение при помощи WairGroup
func ConcurrentSquareSumWG(nums []int) {
	var result int
	
	wg := sync.WaitGroup{}
	mu := sync.Mutex{}

	for _, n := range nums {
		wg.Add(1)        // Добавляем элементы в очередь
		go func(n int) { // Здесь мы прокидываем значеие в функцию
			defer wg.Done()
			mu.Lock() // Происходит блокировка доступа
			result += square(n)

			mu.Unlock()
		}(n)
	}

	wg.Wait()
	fmt.Println(result)

}

func main() {
	nums := []int{2, 4, 6, 8, 10}

	fmt.Println("Starting ConcurrentSquareChannel:")
	ConcurrentSquareSumChannel(nums)
	fmt.Println("Starting ConcurrentSquareWG:")
	ConcurrentSquareSumWG(nums)
}
