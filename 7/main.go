package main

import (
	"fmt"
	"sync"
)

// ConcurrentWriteDataWithMutex выполняет параллельную запись данных с использованием мьютекса
func ConcurrentWriteDataWithMutex() {
	// Количество записей
	writes := 1000

	// Создание хранилища для данных
	storage := make(map[int]int, writes)

	wg := sync.WaitGroup{}
	mu := sync.Mutex{}

	wg.Add(writes)
	for i := 0; i < writes; i++ {
		i := i
		go func() {
			defer wg.Done()

			// Блокировка мьютекса перед обновлением хранилища
			mu.Lock()
			defer mu.Unlock()

			// Запись данных в хранилище
			storage[i] = i
		}()
	}
	// Ожидание завершения всех горутин
	wg.Wait()
	fmt.Println(storage)
}

// ConcurrentWriteDataWithSyncMap выполняет параллельную запись данных с использованием sync.Map
func ConcurrentWriteDataWithSyncMap() {
	writes := 1000

	m := sync.Map{}
	wg := sync.WaitGroup{}

	storage := make(map[int]int)
	
	wg.Add(writes)
	for i := 0; i < writes; i++ {
		i := i 
		go func() {
			defer wg.Done()

			// Запись данных в sync.Map
			m.Store(i, i)
		}()
	}

	wg.Wait()

	// Конвертация данных из sync.Map в обычный map
	m.Range(func(key, value interface{}) bool {
		storage[key.(int)] = value.(int)
		return true
	})

	fmt.Println(storage)
}



func main() {
	ConcurrentWriteDataWithMutex()
	ConcurrentWriteDataWithSyncMap()
}