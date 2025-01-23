package order

import (
	"encoding/json"
	"log"
	"os"

	//"os"
	"time"
)

type Order struct {
	Id        int       `json:"id"`
	Item      string    `json:"item"`
	Quantity  int       `json:"quantity"`
	TimeStamp time.Time `json:"timeStamp"`
}

func (o Order) ToJson() {
	filename := "data/data.json"
	var orders []Order
	file,err := os.ReadFile(filename)
	if err == nil && len(file) > 0{
		err = json.Unmarshal(file,&orders)
		if err != nil{
			log.Println("Error: " , err)
			orders = []Order{}
		}
	}else{
		orders = []Order{}
	}
	orders = append(orders, o)
	jsonData,err := json.MarshalIndent(orders,"","  ")
	if err != nil{
		log.Println("Error : ",err)
		return
	}
	err = os.WriteFile(filename, jsonData,0644)
	if err != nil{
		log.Println("Error: ",err)
		return 
	}

}
