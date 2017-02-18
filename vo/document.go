package vo

type Document struct {
	Id           int    `json:"id"`
	Name         string `json:"name"`
	Price        int    `json:"price"`
	Is_buy       int    `json:"is_buy"`
	Ext          string `json:"ext"`
	Remote_path  string `json:"remote_url,omitempty"`
	Preview_path string `json:"preview_url"`
	Cover_path   string `json:"cover_path"`
	Origin_price int    `json:"origin_price"`
	Buyer_num    int    `json:"buyer_num"`
}
