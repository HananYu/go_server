package models

type User struct {
	Id int `gorm:"primary_key;AUTO_INCREMENT" json:"id"`
	Name string `json:"name"`
	Age int `json:"age"`
	Phone string `json:"phone"`
	Mail string `json:"mail"`
	IdCard string `json:"idCard"`
	Account string `json:"account"`
	Password string `json:"password"`
	Salt string `json:"salt"`
	Avatar string `json:"avatar"`
	CreateDate int `json:"createDate"`
	CreateBy int `json:"createBy"`
	UpdateDate int `json:"updateDate"`
	UpdateBy int `json:"updateBy"`
	IsDel byte `json:"isDel"`

}