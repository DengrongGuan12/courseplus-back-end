package vo

type Video struct {
	Id         int    `json:"id"`
	Video_url  string `json:"video_url"`
	File_size  int64  `json:"file_size"`
	Start_time int64  `json:"start_time"`
	End_time   int64  `json:"end_time"`
}
