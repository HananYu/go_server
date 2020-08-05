package models

/**
分页请求封装
*/

type PageRequest struct {
	PageSize    int `json:"pageSize"`    //每页大小
	CurrentPage int `json:"currentPage"` //当前页数
}
