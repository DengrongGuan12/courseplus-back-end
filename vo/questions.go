package vo

type Questions struct {
	List   []Question `json:"list"`
	Offset int        `json:"offset"`
	Limit  int        `json:"limit"`
	Count  int64      `json:"count"`
}
