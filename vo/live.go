package vo

import (
	"time"
)

type Live struct {
	Id           int          `json:"id"`
	Live_url     string       `json:"live_url"`
	Start_time   time.Time    `json:"start_time"`
	State        string       `json:"state"`
	Period       SimplePeriod `json:"period"`
	Preview_url  string       `json:"preview_url"`
	Complete_url string       `json:"complete_url"`
	Push_url     string       `json:"push_url,omitempty"`
	Video_list   []Video      `json:"video_list"`
}
