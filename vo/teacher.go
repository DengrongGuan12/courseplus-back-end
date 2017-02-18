package vo

type Teacher struct {
	Id           int    `json:"id"`
	User         User   `json:"user"`
	Img_url      string `json:"img_url"`
	Introduction string `json:"introduction"`
	Description  string `json:"description"`
	Education    string `json:"education"`
}
