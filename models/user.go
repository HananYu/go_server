package models

type User struct {
	Id         int    `gorm:"primary_key;AUTO_INCREMENT" json:"id"`
	Name       string `json:"name"`
	Age        int    `json:"age"`
	Phone      string `json:"phone"`
	Mail       string `json:"mail"`
	IdCard     string `json:"id_card"`
	Account    string `json:"account"`
	Password   string `json:"password"`
	Salt       string `json:"salt"`
	Avatar     string `json:"avatar"`
	CreateDate int    `json:"create_date"`
	CreateBy   int    `json:"create_by"`
	UpdateDate int    `json:"update_date"`
	UpdateBy   int    `json:"update_by"`
	IsDel      byte   `json:"is_del"`
}
