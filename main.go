// main.go
package main

import (
	"log"
	"sync"
	"time"

	"github.com/Asladck/ProjectGoRoutines/order"
	"github.com/Asladck/ProjectGoRoutines/processor"
)

func main() {
	
	log.Println("Запуск системы обработки заказов...\n")

	orderChannel := make(chan order.Order, 5) // Буферизированный канал заказов
	quitChannel := make(chan bool)
	var wg sync.WaitGroup
	for i:=0;i<5;i++{
	wg.Add(1)
	go processor.GenerateOrders(orderChannel, quitChannel,wg)
	}
	log.Println("Generating orders...")
	wg.Wait()
	for i:=0;i<5;i++{
	go processor.ProcessOrders(orderChannel, quitChannel)
	}
	// Р`аботаем 10 секунд, потом останавливаемся
	time.Sleep(10 * time.Second)
	quitChannel <- true
	quitChannel <- true // Останавливаем обе горутины
	
	log.Println("Система завершена.")
}
