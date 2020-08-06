package apis

import (
	"gin/config"
	"gin/models"
	"github.com/gin-gonic/gin"
	"github.com/satori/go.uuid"
	"net/http"
	"strconv"
	"strings"
	"time"
)

func UploadFile(c *gin.Context) {
	//单文件
	file, err := c.FormFile("file")
	if err != nil {
		c.JSON(http.StatusOK, models.ReqCode)
		return
	}
	name := file.Filename
	i := strings.LastIndex(name, ".") //含有sub字段的位置
	name = strings.ReplaceAll(uuid.Must(uuid.NewV4(), err).String(), "-", "") + name[i:]
	// 上传文件到指定的路径，保存的文件路径含有文件的名称
	c.SaveUploadedFile(file, config.Save_Path_URL+name)
	c.JSON(http.StatusOK, models.RetunMsgFunc(models.SuccCode, config.Service_URL+name))
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

//获取文章列表
func GetArticleList(c *gin.Context) {
	var page models.PageRequest
	c.BindJSON(&page)
	if page.PageSize == config.Common_ZERO {
		//设置每页大小默认值
		page.PageSize = config.Common_FIVE
	}
	if page.CurrentPage == config.Common_ZERO {
		//设置当前页默认值
		page.CurrentPage = config.Common_ONE
	}
	var arts []models.Article
	config.Db.Table("work_article").Order("create_date desc").Limit(page.PageSize).Offset((page.CurrentPage - config.Common_ONE) * page.PageSize).Find(&arts)
	c.JSON(http.StatusOK, models.RetunMsgFunc(models.SuccCode, arts))
}
