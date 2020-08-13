# go_server
Go项目博客，修改后台java为go

windows打包go部署到linux
SET CGO_ENABLED=0
SET GOOS=linux
SET GOARCH=amd64
go build main.go
