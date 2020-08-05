package apis

import (
	"gin/config"
	"gin/models"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
	"time"
)

func UploadFile(c *gin.Context) {
	//单文件
	file, err := c.FormFile("file")
	if err != nil {
		c.JSON(http.StatusOK, models.ReqCode)
		return
	}
	//fmt.Println(file.Filename)
	// 上传文件到指定的路径，保存的文件路径含有文件的名称
	c.SaveUploadedFile(file, config.Save_Path_URL+file.Filename)
	c.JSON(http.StatusOK, models.RetunMsgFunc(models.SuccCode, config.Service_URL+file.Filename))
}

//保存文件接口
func InserTArticle(c *gin.Context) {
	var article models.Article
	err := c.BindJSON(&article)
	if err != nil {
		c.JSON(http.StatusOK, models.ReqCode)
		return
	}
	if article.Type == config.Common_ZERO || article.Title == "" {

	}
	_, err = strconv.Atoi(c.GetHeader(config.Token_USERID)) // 用户ID
	if err != nil {
		article.CreateBy = 1
	}
	article.CreateDate = int(time.Now().Unix())
	article.ReadNum = config.Common_ZERO
	article.IsDel = config.Common_ZERO

	config.Db.Table("work_article").Create(&article)
	c.JSON(http.StatusOK, models.SuccCode)
}
