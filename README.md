# go_server
Go项目博客，修改后台java为go

修改文件目录
route.go 修改虚拟文件夹路径
common.go 修改文件保存路径和文件访问路径
mysql.go 修改数据库信息
config.js 修改页面访问url

windows打包go部署到linux
SET CGO_ENABLED=0
SET GOOS=linux
SET GOARCH=amd64
go build main.go
