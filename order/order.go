package order

import (
	"encoding/json"
	//"os"
	"time"
)

type Order struct {
	Id        int       `json: "id"`
	Item      string    `json: "item"`
	Quantity  int       `json: "quantity"`
	TimeStamp time.Time `json: "timeStamp"`
}

func (o Order) ToJson() []byte {
	data, err := json.MarshalIndent(o,"","  " )
	if err != nil{
		panic(err)
	}
	return data
	//encoder := 	json.NewEncoder(file)
	// err = encoder.Encode(key)
	// if err != nil{
	// 	panic(err)
	// }
	// file.Close()
}
