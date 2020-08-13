package models

/**
封装返回结构
*/
import (
	"net/http"
)

//--------Code返回码结构体
type Code struct {
	Code int    `json:"code"`
	Msg  string `json:"msg"`
}

//--------定义返回码，以及返回信息
var SuccCode = Code{Code: http.StatusOK, Msg: "success"}
var TokenCode = Code{Code: http.StatusUnauthorized, Msg: "授权已过期"}
var HidenCode = Code{Code: http.StatusUnauthorized, Msg: "访问未授权"}
var AccountCode = Code{Code: http.StatusBadRequest, Msg: "用户已经存在，不允许使用！！！"}
var LoginCode = Code{Code: http.StatusBadRequest, Msg: "用户名或者密码错误！！！"}
var UserCode = Code{Code: http.StatusBadRequest, Msg: "用户不存在！！！"}
var ReqCode = Code{Code: http.StatusBadRequest, Msg: "入参缺失，请补充"}
var SysCode = Code{Code: http.StatusBadRequest, Msg: "系统错误"}

type BaseResult struct {
	Code int         `json:"code"`
	Msg  string      `json:"msg"`
	Data interface{} `json:"data"`
}

//--------对需要返回的信息进行赋值，并以结构体返回
func RetunMsgFunc(code Code, data interface{}) *BaseResult {
	rm := new(BaseResult)
	rm.Code = code.Code
	rm.Msg = code.Msg
	rm.Data = data
	return rm
}
