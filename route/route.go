package route

import (
	. "gin/apis"
	. "gin/utils"
	"github.com/gin-gonic/gin"
	"net/http"
	"strings"
)



//将路由放至在此处
func InitRouter() *gin.Engine {
	router := gin.Default()

	// Use(Authorize())之前的接口，都不用经过身份验证
	router.POST("/api/basic/login", Login)
	router.GET("/user/fU", FindUser)

	//网页跳转
	router.GET("/index", IndexHtml)
	router.GET("/gustbook", GuestBookHtml)



	//以下的接口，都使用Authorize()中间件身份验证
	router.Use(Authorize())

	router.GET("/getH", GetHttp)


	return router
}


func Authorize() gin.HandlerFunc{
	return func(c *gin.Context){
		username := c.Query("username") // 用户名
		ts := c.Query("ts") // 时间戳
		token := c.Query("token") // 访问令牌

		if strings.ToLower(MD5([]byte(username+ts+TokenSalt))) == strings.ToLower(token) {
			// 验证通过，会继续访问下一个中间件
			c.Next()
		} else {
			// 验证不通过，不再调用后续的函数处理
			c.Abort()
			c.JSON(http.StatusUnauthorized,gin.H{"message":"访问未授权"})
			// return可省略, 只要前面执行Abort()就可以让后面的handler函数不再执行
			//return
		}
	}
}