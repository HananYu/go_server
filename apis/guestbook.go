package apis

import (
	"gin/config"
	"gin/models"
	"github.com/gin-gonic/gin"
	"net/http"
	"time"
)

//新增留言
func InsetGuestBook(c *gin.Context) {
	var book models.GuestBook
	err := c.BindJSON(&book)
	if err != nil {
		c.JSON(http.StatusOK, models.ReqCode)
		return
	}
	if book.Aid == 0 || book.NikeName == "" || book.Content == "" {
		c.JSON(http.StatusOK, models.ReqCode)
		return
	}
	//获取请求IP
	//loginip := strings.Split(c.Request.RemoteAddr, ":")[0]
	//fmt.Println(loginip)
	//book.Ip = loginip
	book.CreateTime = int(time.Now().Unix())
	config.Db.Table("work_review").Create(&book)
	c.JSON(http.StatusOK, models.SuccCode)
}
