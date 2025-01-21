package processor

import (
	"math/rand"
	"time"

	"github.com/Asladck/ProjectGoRoutines/order"
)

func GenerateOrders(orderChan chan <- order.Order, boolChan <- chan  bool) {
	defer close(orderChan)
	select{
	case <- boolChan:
		return
	default :
		for i:= 0;i<100;i++{
			orderChan <- order.Order{
				id:       i,
				item:     "Sumka",
				quantity: 2,
				timeStamp: time.Second * 2,
			}
		}
	}
}