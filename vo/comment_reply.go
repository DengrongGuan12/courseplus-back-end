package vo

import "time"

type Comment_reply struct {
	Id          int       `json:"id"`
	Content     string    `json:"content"`
	User        User      `json:"user"`
	Create_time time.Time `json:"create_time"`
}
