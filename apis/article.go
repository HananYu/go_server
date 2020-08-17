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
	c.SaveUploadedFile(file, config.SavePathUrl+name)
	c.JSON(http.StatusOK, models.RetunMsgFunc(models.SuccCode, config.ServiceUrl+name))
}

//新增文章接口
func InserTArticle(c *gin.Context) {
	var article models.Article
	err := c.BindJSON(&article)
	if err != nil {
		c.JSON(http.StatusOK, models.ReqCode)
		return
	}
	if article.Title == "" || article.SmallContent == "" || article.Content == "" {
		c.JSON(http.StatusOK, models.ReqCode)
		return
	}
	userId, bol := c.Get("userId")
	if !bol {
		c.JSON(http.StatusOK, models.UserCode)
		return
	}
	id, _ := userId.(int)
	if id == config.CommonZero {
		c.JSON(http.StatusOK, models.ReqCode)
		return
	}
	article.CreateDate = int(time.Now().Unix())
	//article.ReadNum = config.Common_ZERO
	//article.IsDel = config.Common_ZERO

	config.Db.Table("work_article").Create(&article)
	c.JSON(http.StatusOK, models.SuccCode)
}

//获取文章列表
func GetArticleList(c *gin.Context) {
	var page models.PageRequest
	c.BindJSON(&page)
	if page.PageSize == config.CommonZero {
		//设置每页大小默认值
		page.PageSize = config.CommonFive
	}
	if page.CurrentPage == config.CommonZero {
		//设置当前页默认值
		page.CurrentPage = config.CommonOne
	}
	var arts []models.Article
	config.Db.Table("work_article").Where("is_del=?", config.CommonZero).Order("create_date desc").Limit(page.PageSize).Offset((page.CurrentPage - config.CommonOne) * page.PageSize).Find(&arts)
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
	config.Db.Table("work_article").Order("create_date desc").Where("title LIKE ? or small_content LIKE ? or label like ? ", name, name, name).Find(&arts)
	for i, n := 0, len(arts); i < n; i++ {
		if arts[i].Type == config.CommonZero {
			arts[i].TypeName = "随笔"
		}
		if arts[i].Type == config.CommonOne {
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
	returnMap := make(map[string]interface{}, config.CommonFour)
	var art models.ArticleName
	config.Db.Table("work_article").Where("id = ?", id).First(&art)
	if art.Type == config.CommonZero {
		art.TypeName = "随笔"
	}
	if art.Type == config.CommonOne {
		art.TypeName = "随记"
	}
	returnMap["obj"] = art //当前文章详情
	var lastObj models.ArticleSim
	config.Db.Raw("select id, img,  title from work_article where id =(select id from work_article where is_del = 0 and id < ? order by id desc limit 1)", id).Scan(&lastObj)
	returnMap["lastObj"] = lastObj //上一条文章
	var nextObj models.ArticleSim
	config.Db.Raw("select id, img,  title from work_article where id =(select id from work_article where is_del = 0 and id > ? order by id desc limit 1)", id).Scan(&nextObj)
	returnMap["nextObj"] = nextObj //下一条文章
	var books []models.GuestBook
	config.Db.Table("work_review").Order("create_time desc").Where("a_id = ?", id).Find(&books)
	returnMap["comList"] = books //文章的评论
	//更新访问量+1
	config.Db.Exec("UPDATE work_article SET read_num = read_num + 1 where id = ?", id)
	c.JSON(http.StatusOK, models.RetunMsgFunc(models.SuccCode, returnMap))
}

//仅获取文章详情，用于修改
func GetArticleDetail(c *gin.Context) {
	cid := c.Query("id")
	id, err := strconv.Atoi(cid)
	if err != nil {
		c.JSON(http.StatusOK, models.ReqCode)
		return
	}
	var art models.Article
	config.Db.Table("work_article").Where("id = ?", id).First(&art)
	c.JSON(http.StatusOK, models.RetunMsgFunc(models.SuccCode, art))
}

//更新文章
func UploadArticleDetail(c *gin.Context) {
	var art models.Article
	err := c.BindJSON(&art)
	if err != nil {
		c.JSON(http.StatusOK, models.ReqCode)
		return
	}
	if art.Id == 0 || art.Title == "" || art.SmallContent == "" || art.Content == "" {
		c.JSON(http.StatusOK, models.ReqCode)
		return
	}
	userId, bol := c.Get("userId")
	if !bol {
		c.JSON(http.StatusOK, models.UserCode)
		return
	}
	id, _ := userId.(int)
	if id == config.CommonZero {
		c.JSON(http.StatusOK, models.ReqCode)
		return
	}
	art.UpdateBy = id
	art.UpdateDate = int(time.Now().Unix())
	config.Db.Table("work_article").Save(&art)
	c.JSON(http.StatusOK, models.SuccCode)
}

//删除文章，逻辑删除
func DelArticleByIsDel(c *gin.Context) {
	cid := c.Query("id")
	id, err := strconv.Atoi(cid)
	if err != nil {
		c.JSON(http.StatusOK, models.ReqCode)
		return
	}
	config.Db.Table("work_article").Where("id=?", id).Update("is_del", config.CommonOne)
	c.JSON(http.StatusOK, models.SuccCode)
}
