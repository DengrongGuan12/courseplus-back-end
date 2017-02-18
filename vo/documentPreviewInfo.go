package vo

type DocumentPreviewInfo struct {
	Id           int    `json:"id"`
	Price        int    `json:"-"`
	Preview_path string `json:"preview_url"`
	Remote_path  string `json:"remote_url,omitempty"`
	Is_buy       int    `json:"is_buy"`
	Course_id    int    `json:"-"`
}
