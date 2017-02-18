package vo

import (
	"time"
)

type Period struct {
	Id           int       `json:"id"`
	Name         string    `json:"name"`
	Date         time.Time `json:"date"`
	Price        int       `json:"price"`
	Origin_price int       `json:"origin_price"`
	Content      string    `json:"content"`
	Number       int       `json:"number"`
	Is_buy       int8      `json:"is_buy"`
	Live         Live      `json:"live"`
}
