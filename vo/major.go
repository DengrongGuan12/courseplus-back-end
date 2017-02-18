package vo

type Major struct {
	Id       int      `json:"id"`
	Name     string   `json:"name"`
	Code     string   `json:"code"`
	School   School   `json:"school"`
	Category Category `json:"category"`
}
