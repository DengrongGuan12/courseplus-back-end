package vo

type SimpleLive struct {
	Id       int          `json:"id"`
	Live_url string       `json:"live_url"`
	Period   SimplePeriod `json:"period"`
}
