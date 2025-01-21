// main.go
package main

import (
	"golang.org/x/sync/errgroup"
	"GoRoad/ProjectGoRoutines/order"
)

func main() {
	
	log.Println("Запуск системы обработки заказов...")

	orderChannel := make(chan order.Order, 5) // Буферизированный канал заказов
	quitChannel := make(chan bool)

	// Запускаем генерацию заказов
	go processor.GenerateOrders(orderChannel, quitChannel)

	// Запускаем обработку заказов
	go processor.ProcessOrders(orderChannel, quitChannel)

	// Работаем 10 секунд, потом останавливаемся
	time.Sleep(10 * time.Second)
	quitChannel <- true
	quitChannel <- true // Останавливаем обе горутины

	log.Println("Система завершена.")
}
