package vo

import (
	"time"
)

type Answer struct {
	Id          int           `json:"id"`
	Content     string        `json:"content"`
	Type        string        `json:"type"`
	Create_time time.Time     `json:"create_time"`
	Length      int           `json:"length"`
	Buyer_num   int           `json:"buyer_num"`
	Teacher     SimpleTeacher `json:"teacher"`
}
