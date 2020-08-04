package main

import (
	db "gin/config"
	"gin/route"
)

func main() {
	defer db.Db.Close()
	router := route.InitRouter()
	//http.Handle("/static/", http.StripPrefix("/static/", http.FileServer(http.Dir("static"))))// 启动静态文件服务
	router.LoadHTMLGlob("template/*")
	router.Static("/statics", "./statics")// 启动静态文件服务
	router.Run(":8000")
}