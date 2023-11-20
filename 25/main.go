package main

import (
	"fmt"
	"time"
)

func Sleep(d time.Duration) {
	<-time.After(d) // Здесь происходит блокировка горутины и ожидание сигнала истечения времени
}

func main() {
	fmt.Println("Start func Sleep")
	Sleep(5 * time.Second)
	fmt.Println("Func Sleep succes")
}
