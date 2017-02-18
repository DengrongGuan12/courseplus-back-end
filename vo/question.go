package vo

import "time"

type Question struct {
	Id          int       `json:"id"`
	Content     string    `json:"content"`
	User        *User     `json:"user"`
	Answer      *Answer   `json:"answer,omitempty"`
	Create_time time.Time `json:"create_time"`
	State       string    `json:"state"`
}
