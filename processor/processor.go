package processor

import (
	// "encoding/json"
	// "log"
	"sync"
	"math/rand"
	"os"
	"time"
	"github.com/Asladck/ProjectGoRoutines/order"
)


func GenerateOrders(orderChan chan <- order.Order, boolChan <- chan  bool,wg sync.WaitGroup) {
	list := []string{"Bag","Car","Cow","Dota2"}
	n:=0
	select{
	
	case <- boolChan:
		return
	default :
		n++
			orderChan <- order.Order{
				Id:       n,
				Item:     list[int(rand.Float64()*4)],
				Quantity: int(rand.Float64()*20),
				TimeStamp: time.Now(),
		}
	}
	wg.Done()
}
func ProcessOrders(orderChan <-chan order.Order,boolChan <- chan  bool){
		select{
		case <- boolChan:
			return 
		case v := <- orderChan:
			os.WriteFile("data/data.json",v.ToJson(),0644)
		}
	}