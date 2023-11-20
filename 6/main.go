package main

import (
	"context"
	"fmt"
	"os"
	"os/signal"
	"sync"
	"syscall"
	"time"
)

var wg sync.WaitGroup

// CancelGoroutineWithChannel демонстрирует отмену горутины с использованием канала.
func CancelGoroutineWithChannel() {
	// Создаем канал для уведомления о завершении работы горутины.
	quit := make(chan bool)
	// Добавляем горутину в группу ожидания.
	wg.Add(1)
	// Запускаем горутину.
	go func() {
		defer wg.Done()
		for {
			select {
			case <-quit:
				fmt.Println("Channel finished working")
				return
			default:
				fmt.Println("Channel working")
				time.Sleep(time.Second)
			}
		}
	}()
	// Ждем 3 секунды.
	time.Sleep(time.Second * 3)
	// Закрываем канал для завершения работы горутины.
	quit <- true
	wg.Wait()

}

// CancelGoroutineUsingContext демонстрирует отмену горутины с использованием контекста.
func CancelGoroutineUsingContext() {
	// Создаем контекст с функцией отмены.
	ctx, cancel := context.WithCancel(context.Background())
	// defer cancel()

	wg.Add(1)
	go func(ctx context.Context) {
		defer wg.Done()
		for {
			select {
			case <-ctx.Done():
				fmt.Println("Context cancelled")
				return
			default:
				fmt.Println("Gorutina through context still works")
				time.Sleep(time.Second)

			}
		}
	}(ctx)

	time.Sleep(time.Second * 3)
	// Отменяем контекст для завершения работы горутины.
	cancel()
	wg.Wait()

}

// CancelGoroutineUsingSignals демонстрирует отмену горутины с использованием сигналов ОС
func CancelGoroutineUsingSignals() {
	// Создаем канал для получения сигналов от ОС
	osSignal := make(chan os.Signal, 1)

	// Регистрируем сигналы для получения уведомлений.
	signal.Notify(osSignal, os.Interrupt, syscall.SIGTERM)

	// Создаем канал для уведомления о завершении работы горутины.
	doneSignal := make(chan struct{})

	go func() {
		<- osSignal
		close(doneSignal)
		fmt.Println("Goroutine signals completed")
	}()

	fmt.Println("Ctrl+C stopped goroutine")

	// Ждем завершения работы горутины по сигналу от ОС
	<- doneSignal

}

func main() {

	fmt.Println("Using context")
	CancelGoroutineUsingContext()

	fmt.Println("Using channel")
	CancelGoroutineWithChannel()

	fmt.Println("Using signals")
	CancelGoroutineUsingSignals()
}
