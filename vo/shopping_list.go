package vo

type ShoppingList struct {
	Period_list    []SimplePeriod   `json:"period_list"`
	Document_list  []SimpleDocument `json:"document_list"`
	Question_count int              `json:"question_count"`
	Price          int              `json:"price"`
}
