package vo

import (
	"time"
)

type SimplePeriod struct {
	Id     int       `json:"id"`
	Name   string    `json:"name"`
	Number int       `json:"number"`
	Date   time.Time `json:"date"`
	Price  int       `json:"price"`
	Is_buy int8      `json:"is_buy"`
}
