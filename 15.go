package main

import (
	"fmt"
	"time"
)

func main() {
	// Длительность работы программы
	N := 5 * time.Second

	// Создание канала
	ch := make(chan int)

	// Запуск горутины для записи в канал
	go func() {
		counter := 0
		for {
			time.Sleep(100 * time.Millisecond)
			ch <- counter
			counter++
		}
	}()

	// Таймер для завершения программы
	go func() {
		time.Sleep(N) // Используем time.Sleep вместо time.After
		close(ch)     // Закрываем канал для завершения программы
	}()

	// Чтение из канала
	for val := range ch {
		fmt.Println("Received:", val)
	}
	fmt.Println("Program completed.")
}
