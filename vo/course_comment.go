package vo

import "time"

type Course_comment struct {
	Id          int       `json:"id"`
	Content     string    `json:"content"`
	Create_time time.Time `json:"create_time"`
	Star        float32   `json:"star"`
	Order       Order     `json:"-"`
	Reply_count int       `json:"reply_count"`
	User        User      `json:"user"`
	Order_info  string    `json:"order_info,omitempty"`
}
