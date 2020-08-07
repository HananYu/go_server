package models

type GuestBook struct {
	Id         int    ` json:"id" gorm:"primary_key;AUTO_INCREMENT"`
	Aid        int    ` json:"aId" gorm:"column:a_id"` //后面这个为数据库列名
	NikeName   string `json:"nikeName" gorm:"column:nike_name"`
	Content    string `json:"content"`
	CreateTime int    `json:"createTime" gorm:"column:create_time"`
	Email      string `json:"email"`
	Phone      string `json:"phone"`
	Ip         string `json:"ip"`
}
