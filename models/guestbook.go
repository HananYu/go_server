package models

type GuestBook struct {
	Id         int    `json:"id"`
	Aid        int    `json:"a_id"`
	NikeName   string `json:"nike_name"`
	Content    string `json:"content"`
	CreateTime int    `json:"create_time"`
	Email      string `json:"email"`
	Phone      string `json:"phone"`
	Ip         string `json:"ip"`
}
