package models

type Article struct {
	Id           int    `json:"id"`
	CreateDate   int    `json:"create_date"`
	CreateBy     int    `json:"create_by"`
	UpdateDate   int    `json:"update_date"`
	UpdateBy     int    `json:"update_by"`
	Title        string `json:"title"`
	Type         int    `json:"type"`
	SmallContent string `json:"small_content"`
	Content      string `json:"content"`
	Img          string `json:"img"`
	ReadNum      int    `json:"read_num"`
	IsDel        int    `json:"is_del"`
}

type ArticleName struct {
	Id           int    `json:"id"`
	CreateDate   int    `json:"create_date"`
	CreateBy     int    `json:"create_by"`
	UpdateDate   int    `json:"update_date"`
	UpdateBy     int    `json:"update_by"`
	Title        string `json:"title"`
	Type         int    `json:"type"`
	SmallContent string `json:"small_content"`
	Content      string `json:"content"`
	Img          string `json:"img"`
	ReadNum      int    `json:"read_num"`
	IsDel        int    `json:"is_del"`
	TypeName     string `json:"typeName"`
}

type ArticleSim struct {
	Id    int    `json:"id"`
	Img   string `json:"img"`
	Title string `json:"title"`
}
