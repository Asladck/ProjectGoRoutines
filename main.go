// main.go
package main

import (
	"time"

	"github.com/Asladck/ProjectGoRoutines/order"
	"github.com/Asladck/ProjectGoRoutines/processor"
)

func main() {
	
	print("Запуск системы обработки заказов...\n")

	orderChannel := make(chan order.Order, 5) // Буферизированный канал заказов
	quitChannel := make(chan bool)

	// Запускаем генерацию заказов
	go processor.GenerateOrders(orderChannel, quitChannel)
	v := order.Order{
		id : 2,
		item : "",
		quantity : 3,
		timeStamp : time.Second,
	}
	print(v)
	// Запускаем обработку заказов
	go processor.ProcessOrders(orderChannel, quitChannel)

	// Работаем 10 секунд, потом останавливаемся
	time.Sleep(10 * time.Second)
	quitChannel <- true
	quitChannel <- true // Останавливаем обе горутины

	log.Println("Система завершена.")
}
