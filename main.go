package main

import (
	db "gin/config"
	"gin/route"
)

func main() {
	// 禁用控制台颜色
	//gin.DisableConsoleColor()
	// 创建记录日志的文件
	//f, _ := os.Create("gin.log")
	//gin.DefaultWriter = io.MultiWriter(f)
	//// 如果需要将日志同时写入文件和控制台，请使用以下代码
	//gin.DefaultWriter = io.MultiWriter(f, os.Stdout)

	defer db.Db.Close()          //defer延时关闭，在当前方法完成后执行
	router := route.InitRouter() //初始化router，以及路由管理
	router.Run(":8000")          //设置端口号
}
