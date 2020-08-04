package models

type Article struct {
	Id int `json:"id"`
	Create_date int `json:"create_date"`
	create_by int `json:"create_by"`
	update_date int `json:"update_date"`
	update_by int `json:"update_by"`
	title int `json:"title"`
	Type int `json:"type"`
	small_content int `json:"small_content"`
	content int `json:"content"`
	img int `json:"img"`
	read_num int `json:"read_num"`
	is_del int `json:"is_del"`
}
