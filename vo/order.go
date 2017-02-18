package vo

import (
	"time"
)

type Order struct {
	Id          int		`json:"id"`
	Orderno     string	`json:"orderno"`
	User        *User 	`json:"user"`
	Type        string	`json:"type"`
	Object_id   int		`json:"object_id"`
	Charge_id   string	`json:"charge_id"`
	Description string	`json:"description"`
	Price	    int		`json:"price"`
	Commented   int8      	`json:"commented"`
	Pay_time    *time.Time 	`json:"pay_time"`
	Create_time time.Time 	`json:"create_time"`
}
