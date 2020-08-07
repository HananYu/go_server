package models

type Article struct {
	Id           int    `json:"id" gorm:"primary_key;AUTO_INCREMENT"`
	CreateDate   int    `json:"createDate" gorm:"column:create_date"`
	CreateBy     int    `json:"createBy" gorm:"column:create_by"`
	UpdateDate   int    `json:"updateDate" gorm:"column:update_date"`
	UpdateBy     int    `json:"updateBy" gorm:"column:update_by"`
	Title        string `json:"title"`
	Type         int    `json:"type"`
	SmallContent string `json:"smallContent" gorm:"column:small_content"`
	Content      string `json:"content"`
	Img          string `json:"img"`
	ReadNum      int    `json:"readNum" gorm:"column:read_num"`
	IsDel        int    `json:"isDel" gorm:"column:is_del"`
}

type ArticleName struct {
	Id           int    `json:"id" gorm:"primary_key;AUTO_INCREMENT"`
	CreateDate   int    `json:"createDate" gorm:"column:create_date"`
	CreateBy     int    `json:"createBy" gorm:"column:create_by"`
	UpdateDate   int    `json:"updateDate" gorm:"column:update_date"`
	UpdateBy     int    `json:"updateBy" gorm:"column:update_by"`
	Title        string `json:"title"`
	Type         int    `json:"type"`
	SmallContent string `json:"smallContent" gorm:"column:small_content"`
	Content      string `json:"content"`
	Img          string `json:"img"`
	ReadNum      int    `json:"readNum" gorm:"column:read_num"`
	IsDel        int    `json:"isDel" gorm:"column:is_del"`
	TypeName     string `json:"typeName"`
}

type ArticleSim struct {
	Id    int    `json:"id"`
	Img   string `json:"img"`
	Title string `json:"title"`
}
