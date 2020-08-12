package route

import (
	"fmt"
	. "gin/apis"
	"gin/models"
	. "gin/utils"
	"github.com/gin-gonic/gin"
	"net/http"
	"strings"
)

//将路由放至在此处
func InitRouter() *gin.Engine {
	router := gin.Default()

	// Use(Authorize())之前的接口，都不用经过身份验证
	router.POST("/api/basic/logon", InsertUser) //注册接口
	router.POST("/api/basic/login", Login)      //登陆接口

	//文章接口
	router.POST("/api/basic/upload", UploadFile)         //上传文件接口
	router.POST("/api/article/add", InserTArticle)       //保存文章接口
	router.POST("/api/article/get", GetArticleList)      //获取文章列表
	router.GET("/api/article/search", SearchArticleList) //搜索文章列表
	router.GET("/api/article/detail", DetailArticleList) //获取文章详情

	//留言接口
	router.POST("/api/guestbook/add", InsetGuestBook) //插入留言
	router.GET("/api/guestbook/get", GetGuestBooks)   //获取留言
	//归档页面接口
	router.POST("/api/record/get", GetArticleRecords) //插入留言

	router.GET("/user/fU", FindUser)

	//网页跳转
	router.GET("/", IndexHtml)
	router.GET("/login", LoginHtml)
	router.GET("/index", IndexHtml)
	router.GET("/gustbook", GuestBookHtml)
	router.GET("/article", ArticleHtml)
	router.GET("/search", SearchHtml)
	router.GET("/archives", ArchivesHtml)
	router.GET("/detail", DetailHtml)
	router.GET("/link", LinkHtml)
	router.GET("/update", UpdateHtml)

	//静态资源，放置在拦截器之前不会对静态资源进行拦截
	router.LoadHTMLGlob("statics/template/*")
	router.Static("statics", "./statics").Static("file", "F:/Temp/") // 启动静态文件服务
	//router.Static("statics", "./statics").Static("file", "/home/temp/") // 启动静态文件服务

	//http.Handle("/statics/", http.StripPrefix("/statics/", http.FileServer(http.Dir("statics"))))// 启动静态文件服务

	//以下的接口，都使用Authorize()中间件身份验证
	router.Use(Authorize())
	router.GET("/api/getH", GetHttp) //用于校验用户token

	return router
}

func Authorize() gin.HandlerFunc {
	return func(c *gin.Context) {
		userId := c.GetHeader("userId") // 用户ID
		salt := c.GetHeader("salt")     // 盐
		token := c.GetHeader("token")   // 访问令牌
		fmt.Println("请求token----------" + strings.ToLower(MD5([]byte(string(userId)+salt+TokenSalt))))
		fmt.Println("存储token----------" + strings.ToLower(token))
		if strings.ToLower(MD5([]byte(userId+salt+TokenSalt))) == strings.ToLower(token) {
			// 验证通过，会继续访问下一个中间件
			c.Next()
		} else {
			// 验证不通过，不再调用后续的函数处理
			c.Abort()
			c.JSON(http.StatusOK, models.HidenCode)
		}
	}
}
