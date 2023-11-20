package main

import (
	"context"
	"fmt"
	"sync"
	"time"
)

// Функция, которая отправляет значения в канал ch
func writer(ctx context.Context, ch chan<- int) {
	for i := 0; true; i++ {
		select {
			// Проверка, если контекст завершен
		case <-ctx.Done():
			// Закрытие канала и завершение функции
			close(ch)
			return
		default:
			// Отправка значения в канал ch
			ch <- i
			time.Sleep(time.Millisecond * 250)

		}
	}
}

// Функция, которая читает значения из канала ch
func reader(ch <-chan int, wg *sync.WaitGroup) {
	// Итерация по значениям в канале
	for num := range ch {
		// Вывод числа в консоль
		fmt.Printf("%d\n", num)
	}
	// Уменьшение счетчика WaitGroup
	wg.Done()
}

func main() {
	// Создание WaitGroup для ожидания завершения горутин
	wg := &sync.WaitGroup{}
	// Создание контекста с таймаутом в 5 секунд
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*5)
	defer cancel()

	// Создание канала с буфером в 1 элемент
	ch := make(chan int, 1)

	// Увеличение счетчика WaitGroup на 1
	wg.Add(1)
	// Запуск горутины для записи значений в канал
	go writer(ctx, ch)
	// Запуск горутины для чтения значений из канала
	go reader(ch, wg)
	// Ожидание завершения всех горутин
	wg.Wait()

}
