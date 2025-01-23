// main.go
package main

import (
	"context"
	"fmt"
	"log"
	"sync"
	"time"

	"github.com/Asladck/ProjectGoRoutines/order"
	"github.com/Asladck/ProjectGoRoutines/processor"
)

func main() {
	fmt.Println("Запуск системы обработки заказов...\n")

	orderChannel := make(chan order.Order,10)// Буферизированный канал заказов
	ctx,cancel := context.WithCancel(context.Background())
	var wg sync.WaitGroup
	for i:=0;i<3;i++{
		wg.Add(1)
			go processor.GenerateOrders(ctx, orderChannel,&wg)		
	}
	log.Println("Generating orders...")
	for i:=0;i<3;i++{
		wg.Add(1)
		
			go processor.ProcessOrders(ctx, orderChannel,&wg)	

	}
	// Р`аботаем 10 секунд, потом останавливаемся
	time.Sleep(5 * time.Second)
	cancel()
	wg.Wait()
	close(orderChannel)
	log.Println("Система завершена.")
}
