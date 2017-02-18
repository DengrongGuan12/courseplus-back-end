package vo

type Category struct {
	Id   int    `json:"id"`
	Name string `json:"name,omitempty"`
	Code string `json:"code,omitempty"`
}
