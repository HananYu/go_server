package models

type GuestBook struct {
	Id         int    ` json:"id" gorm:"primary_key;AUTO_INCREMENT"`
	Aid        int    ` json:"a_id" gorm:"column:a_id"`
	NikeName   string `json:"nike_name"`
	Content    string `json:"content"`
	CreateTime int    `json:"create_time"`
	Email      string `json:"email"`
	Phone      string `json:"phone"`
	Ip         string `json:"ip"`
}
