package order

import (
	"encoding/json"
	"time"
)

type Order struct {
	id        int       `json: "id"`
	item      string    `json: "item"`
	quantity  int       `json: "quantity"`
	timeStamp time.Time `json: "timeStamp"`
}

func (o Order) toJson() string {
	data, _ := json.Marshal(o)
	return string(data)
}
