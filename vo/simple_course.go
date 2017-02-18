package vo

type SimpleCourse struct {
	Id      int           `json:"id"`
	Teacher SimpleTeacher `json:"teacher"`
}
