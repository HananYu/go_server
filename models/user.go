package models

type User struct {
	Id         int    `gorm:"primary_key;AUTO_INCREMENT" json:"id"`
	Name       string `json:"name"`
	Age        int    `json:"age"`
	Phone      string `json:"phone"`
	Mail       string `json:"mail"`
	IdCard     string `json:"idCard" gorm:"column:id_card"`
	Account    string `json:"account"`
	Password   string `json:"password"`
	Salt       string `json:"salt"`
	Avatar     string `json:"avatar"`
	CreateDate int    `json:"createDate" gorm:"column:create_date"`
	CreateBy   int    `json:"createBy" gorm:"column:create_by"`
	UpdateDate int    `json:"updateDate" gorm:"column:update_date"`
	UpdateBy   int    `json:"updateBy" gorm:"column:update_by"`
	IsDel      byte   `json:"isDel" gorm:"column:is_del"`
}
