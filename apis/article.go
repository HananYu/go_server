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

//上传文件
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

//搜索文章列表
func SearchArticleList(c *gin.Context) {
	name := c.Query("s")
	if name == "" {
		c.JSON(http.StatusOK, models.ReqCode)
		return
	}
	name = "%" + name + "%"
	var arts []models.ArticleName
	config.Db.Table("work_article").Order("create_date desc").Where("title LIKE ? or small_content LIKE ? ", name, name).Find(&arts)
	for i, n := 0, len(arts); i < n; i++ {
		if arts[i].Type == config.Common_ZERO {
			arts[i].TypeName = "随笔"
		}
		if arts[i].Type == config.Common_ONE {
			arts[i].TypeName = "随记"
		}
	}
	c.JSON(http.StatusOK, models.RetunMsgFunc(models.SuccCode, arts))
}

//获取文章详情
func DetailArticleList(c *gin.Context) {
	cid := c.Query("id")
	id, err := strconv.Atoi(cid)
	if err != nil {
		c.JSON(http.StatusOK, models.ReqCode)
		return
	}
	returnMap := make(map[string]interface{}, config.Common_FOUR)
	var art models.ArticleName
	config.Db.Table("work_article").Where("id = ?", id).First(&art)
	returnMap["obj"] = art //当前文章详情
	var lastObj models.ArticleSim
	config.Db.Raw("select id, img,  title from work_article where id =(select id from work_article where id < ? order by id desc limit 1)", id).Scan(&lastObj)
	returnMap["lastObj"] = lastObj //上一条文章
	var nextObj models.ArticleSim
	config.Db.Raw("select id, img,  title from work_article where id =(select id from work_article where id > ? order by id desc limit 1)", id).Scan(&nextObj)
	returnMap["nextObj"] = nextObj //下一条文章
	var books []models.GuestBook
	config.Db.Table("work_review").Order("create_time desc").Where("a_id = ?", id).Find(&books)
	returnMap["comList"] = books //文章的评论
	c.JSON(http.StatusOK, models.RetunMsgFunc(models.SuccCode, returnMap))
}
