package vo

type User struct {
	Id       int    `json:"id"`
	Nickname string `json:"name"`
	Img_url  string `json:"img_url"`
	Account  string `json:"account"`
	Gender   int8   `json:"gender"`
}
