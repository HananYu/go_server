package apis

import (
	"gin/config"
	"gin/models"
	"gin/utils"
	"github.com/gin-gonic/gin"
	"github.com/satori/go.uuid"
	"net/http"
	"os"
	"strconv"
	"strings"
	"time"
)

//上传文件
func UploadFile(c *gin.Context) {
	file, err := c.FormFile(config.CommonFile)
	if err != nil {
		c.JSON(http.StatusOK, models.ReqCode)
		return
	}
	name := file.Filename
	i := strings.LastIndex(name, ".") //含有sub字段的位置
	name = strings.ReplaceAll(uuid.Must(uuid.NewV4(), err).String(), "-", config.CommonNull) + name[i:]
	// 上传文件到指定的路径，保存的文件路径含有文件的名称
	_ = c.SaveUploadedFile(file, config.SavePathUrl+name)
	c.JSON(http.StatusOK, models.RetunMsgFunc(models.SuccCode, config.ServiceUrl+name))
}

//上传图片，用于上传封面，需要对图片进行压缩
func UploadFileCompress(c *gin.Context) {
	file, err := c.FormFile(config.CommonFile)
	if err != nil {
		c.JSON(http.StatusOK, models.ReqCode)
		return
	}
	name := file.Filename
	i := strings.LastIndex(name, ".") //含有sub字段的位置
	name = strings.ReplaceAll(uuid.Must(uuid.NewV4(), err).String(), "-", config.CommonNull) + name[i:]
	// 上传文件到指定的路径，保存的文件路径含有文件的名称
	_ = c.SaveUploadedFile(file, config.SavePathUrl+name)
	url := utils.ChangeImage(config.SavePathUrl + name)
	if url == config.CommonNull {
		_ = os.Remove(config.SavePathUrl + name)
		c.JSON(http.StatusOK, models.SysCode)
		return
	}
	c.JSON(http.StatusOK, models.RetunMsgFunc(models.SuccCode, url))
}

//新增文章接口
func InserTArticle(c *gin.Context) {
	var article models.Article
	err := c.BindJSON(&article)
	if err != nil {
		c.JSON(http.StatusOK, models.ReqCode)
		return
	}
	if article.Title == config.CommonNull || article.SmallContent == config.CommonNull || article.Content == config.CommonNull {
		c.JSON(http.StatusOK, models.ReqCode)
		return
	}
	userId, bol := c.Get(config.TokenUSERID)
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
	if article.Type == config.CommonTwo {
		//获取内容里面的所有图片地址，最多取四张，包括封面一张，使用逗号进行分隔
		var list []string
		list = utils.GetImageByString(article.Content, list)
		if len(list) >= config.CommonZero {
			for i, str := range list {
				if i > config.CommonThree {
					break
				}
				s := utils.ChangeImageBySize(config.SavePathUrl + str[strings.Index(str, config.ServiceUrl)+len(config.ServiceUrl):])
				article.Img += config.CommonComma + s
			}
		}
	}
	config.Db.Table(config.DataTableArticle).Create(&article)
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
	var arts []models.ArticleList
	config.Db.Table(config.DataTableArticle).Where("is_del=?", config.CommonZero).Order("create_date desc").Limit(page.PageSize).Offset((page.CurrentPage - config.CommonOne) * page.PageSize).Find(&arts)
	c.JSON(http.StatusOK, models.RetunMsgFunc(models.SuccCode, arts))
}

//搜索文章列表
func SearchArticleList(c *gin.Context) {
	name := c.Query(config.CommonName)
	if name == config.CommonNull {
		c.JSON(http.StatusOK, models.ReqCode)
		return
	}
	name = "%" + name + "%"
	var arts []models.ArticleName
	config.Db.Table(config.DataTableArticle).Order("create_date desc").Where("title LIKE ? or small_content LIKE ? or label like ? ", name, name, name).Find(&arts)
	for i, n := 0, len(arts); i < n; i++ {
		if arts[i].Type == config.CommonZero {
			arts[i].TypeName = config.CommonArticleTypeZero
		}
		if arts[i].Type == config.CommonOne {
			arts[i].TypeName = config.CommonArticleTypeOne
		}
		if arts[i].Type == config.CommonTwo {
			arts[i].TypeName = config.CommonArticleTypeTwo
		}
	}
	c.JSON(http.StatusOK, models.RetunMsgFunc(models.SuccCode, arts))
}

//获取文章详情
func DetailArticleList(c *gin.Context) {
	cid := c.Query(config.CommonId)
	id, err := strconv.Atoi(cid)
	if err != nil {
		c.JSON(http.StatusOK, models.ReqCode)
		return
	}
	returnMap := make(map[string]interface{}, config.CommonFour)
	var art models.ArticleName
	config.Db.Table(config.DataTableArticle).Where("id = ?", id).First(&art)
	if art.Type == config.CommonZero {
		art.TypeName = config.CommonArticleTypeZero
	}
	if art.Type == config.CommonOne {
		art.TypeName = config.CommonArticleTypeOne
	}
	if art.Type == config.CommonTwo {
		art.TypeName = config.CommonArticleTypeTwo
	}
	returnMap["obj"] = art //当前文章详情
	var lastObj models.ArticleSim
	config.Db.Raw("select id, img,  title from work_article where id =(select id from work_article where is_del = 0 and id < ? order by id desc limit 1)", id).Scan(&lastObj)
	returnMap["lastObj"] = lastObj //上一条文章
	var nextObj models.ArticleSim
	config.Db.Raw("select id, img,  title from work_article where id =(select id from work_article where is_del = 0 and id > ? order by id desc limit 1)", id).Scan(&nextObj)
	returnMap["nextObj"] = nextObj //下一条文章
	var books []models.GuestBook
	config.Db.Table(config.DataTableReview).Order("create_time desc").Where("a_id = ?", id).Find(&books)
	returnMap["comList"] = books //文章的评论
	//更新访问量+1
	config.Db.Exec("UPDATE work_article SET read_num = read_num + 1 where id = ?", id)
	c.JSON(http.StatusOK, models.RetunMsgFunc(models.SuccCode, returnMap))
}

//仅获取文章详情，用于修改
func GetArticleDetail(c *gin.Context) {
	cid := c.Query(config.CommonId)
	id, err := strconv.Atoi(cid)
	if err != nil {
		c.JSON(http.StatusOK, models.ReqCode)
		return
	}
	var art models.Article
	config.Db.Table(config.DataTableArticle).Where("id = ?", id).First(&art)
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
	if art.Id == config.CommonZero || art.Title == config.CommonNull || art.SmallContent == config.CommonNull || art.Content == config.CommonNull {
		c.JSON(http.StatusOK, models.ReqCode)
		return
	}
	userId, bol := c.Get(config.TokenUSERID)
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
	config.Db.Table(config.DataTableArticle).Save(&art)
	c.JSON(http.StatusOK, models.SuccCode)
}

//删除文章，逻辑删除
func DelArticleByIsDel(c *gin.Context) {
	cid := c.Query(config.CommonId)
	id, err := strconv.Atoi(cid)
	if err != nil {
		c.JSON(http.StatusOK, models.ReqCode)
		return
	}
	config.Db.Table(config.DataTableArticle).Where("id=?", id).Update("is_del", config.CommonOne)
	c.JSON(http.StatusOK, models.SuccCode)
}
