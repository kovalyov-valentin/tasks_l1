package main

import (
	"context"
	"fmt"
	"math/rand"
	"os"
	"os/signal"
	"sync"
	"syscall"
)

func main() {
	// Запрашиваем у пользователя количество воркеров
	var countWorkers int
	fmt.Println("Select the number of workers")
	fmt.Fscan(os.Stdin, &countWorkers)

	// Создаем канал для передачи работы в воркеры
	workChannel := make(chan int)
	defer close(workChannel) // Закрытие канала после использования

	// Создаем канал для сигналов завершения программы
	exitSignal := make(chan os.Signal, 1)
	defer close(exitSignal)
	signal.Notify(exitSignal, syscall.SIGINT, syscall.SIGTERM) // Оповещение о сигналах завершения

	// Создаем контекст и функцию для отмены работы воркеров
	ctx, cancel := context.WithCancel(context.Background())

	var wg sync.WaitGroup

	// Запускаем указанное количество воркеров
	for i := 0; i < countWorkers; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			for {
				select {
				case data := <-workChannel:
					fmt.Fprintln(os.Stdout, data)
				case <-ctx.Done():
					fmt.Println("Worker finished")
					return
				}
			}
		}()
	}

	// В цикле передаем случайные данные в канал работы воркеров
	for {
		select {
		case workChannel <- rand.Intn(100):
			
		case <-exitSignal:
			cancel()  // Отменяем контекст, что приведет к завершению работы воркеров
			wg.Wait() // Ожидаем завершения всех воркеров
			fmt.Println("Finish")
			return
		}
	}
}
