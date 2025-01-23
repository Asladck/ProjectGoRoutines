package processor

import (
	// "encoding/json"
	// "log"
	"context"
	"math/rand"
	"sync"
	"time"

	"github.com/Asladck/ProjectGoRoutines/order"
)


func GenerateOrders(ctx context.Context, orderChan chan <- order.Order,wg *sync.WaitGroup) {
	defer wg.Done()
	list := []string{"Bag","Car","Cow","Dota2"}
	idCounter :=0 
	for{
	idCounter++
	select{
	case <- ctx.Done():
		return
	default :
			orderChan <- order.Order{
				Id:       idCounter,
				Item:     list[rand.Intn(len(list))],
				Quantity: rand.Intn(20)+1,
				TimeStamp: time.Now(),
		}
		time.Sleep(1000 * time.Millisecond)
	}
	}
}
func ProcessOrders(ctx context.Context,orderChan <-chan order.Order,wg *sync.WaitGroup){
	var mu sync.Mutex
	defer wg.Done()
	for{	
	select{
		case <- ctx.Done():
			return 
		case v,ok := <- orderChan:
			if !ok{
				return
			}
			mu.Lock()
			v.ToJson()
			mu.Unlock()
		}
	}
}