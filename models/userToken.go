package models

type UserToken struct {
	Id         int    `json:"id" gorm:"primary_key;AUTO_INCREMENT"`
	UserId     int    `json:"userId" gorm:"column:user_id"`
	Token      string `json:"token"`
	CreateDate int    `json:"createDate" gorm:"column:create_date"`
	EndDate    int    `json:"endDate" gorm:"column:end_date"`
}
