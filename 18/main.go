package main

import (
	"fmt"
	"sync"
)

type CounterWithMutex struct {
	count int
	sync.Mutex
}

func (c *CounterWithMutex) Increment(wg *sync.WaitGroup) {
	defer wg.Done()
	c.Lock()
	defer c.Unlock()
	c.count++
}


func main() {
	wg := sync.WaitGroup{}
	n := 15 // Количество горутин

	muCounter := &CounterWithMutex{count: 0}

	for i := 0; i < n; i++ { // Запуск горутин, инкрементрирующих счетчик
		wg.Add(1)
		go func() {
			muCounter.Increment(&wg)
		}()
	}
	wg.Wait()
	fmt.Println("Value counter:", muCounter.count)
}
