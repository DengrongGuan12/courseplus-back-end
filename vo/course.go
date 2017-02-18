package vo

type Courses struct {
	Offset int      `json:"offset"`
	Limit  int      `json:"limit"`
	Count  int64    `json:"count"`
	List   []Course `json:"list"`
}

type Course struct {
	Id            int            `json:"id"`
	Name          string         `json:"name"`
	Code          string         `json:"code"`
	Star          float32        `json:"star"`
	Buyer_num     int            `json:"buyer_num"`
	Img_url       string         `json:"img_url"`
	Teacher       Teacher        `json:"teacher"`
	Major         Major          `json:"major"`
	Subject       Subject        `json:"subject"`
	Origin_price  int            `json:"origin_price"`
	Current_price int            `json:"current_price"`
	Is_buy        int8           `json:"is_buy"`
	Period_num    int64          `json:"period_num"`
	Doc_num       int64          `json:"doc_num"`
	Free_ask_num  int            `json:"free_ask_num"`
	Other_courses []SimpleCourse `json:"other_courses,omitempty"`
}
