package route

import (
	. "gin/apis"
	"gin/config"
	"gin/models"
	"github.com/gin-gonic/gin"
	"net/http"
	"time"
)

//将路由放至在此处
func InitRouter() *gin.Engine {
	router := gin.Default()

	// Use(Authorize())之前的接口，都不用经过身份验证
	router.POST("/api/basic/logon", InsertUser)       //注册接口
	router.POST("/api/basic/login", Login)            //登陆接口
	router.GET("/api/basic/account", FindByAccount)   //判断账号是否存在
	router.POST("/api/basic/changPw", ChangePassword) //修改用户密码，传入account password

	//文章接口
	router.POST("/api/basic/upload", UploadFile)         //上传文件接口
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
	//router.GET("/update", UpdateHtml)

	//静态资源，放置在拦截器之前不会对静态资源进行拦截
	router.LoadHTMLGlob("statics/template/*")
	router.Static("statics", "./statics").Static("file", "F:/Temp/") // 启动静态文件服务
	//router.Static("statics", "./statics").Static("file", "/home/temp/") // 启动静态文件服务

	//以下的接口，都使用Authorize()中间件身份验证
	router.Use(Authorize())
	router.GET("/api/getH", GetHttp) //用于校验用户token

	//文章接口
	router.POST("/api/article/add", InserTArticle) //保存文章接口

	return router
}

func Authorize() gin.HandlerFunc {
	return func(c *gin.Context) {
		token := c.GetHeader("token") // 访问令牌

		var ut models.UserToken
		config.Db.Table("sys_user_token").Where("token = ?", token).First(&ut)
		if ut.UserId == config.CommonZero {
			// 验证不通过，不再调用后续的函数处理
			c.Abort()
			c.JSON(http.StatusOK, models.HidenCode)
			return
		}
		if ut.EndDate < int(time.Now().Unix()) {
			//通过过期时间来判断是不是存在这个token或者过期
			config.Db.Table("sys_user_token").Delete(&ut) //删除过期记录
			c.Abort()
			c.JSON(http.StatusOK, models.TokenCode)
			return
		}
		c.Set("userId", ut.UserId)
		c.Next()
	}
}
