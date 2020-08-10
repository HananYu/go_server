package apis

import (
	"gin/config"
	"gin/models"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
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
	if book.NikeName == "" || book.Content == "" {
		c.JSON(http.StatusOK, models.ReqCode)
		return
	}
	book.Ip = c.ClientIP()
	book.CreateTime = int(time.Now().Unix())
	config.Db.Table("work_review").Create(&book)
	c.JSON(http.StatusOK, models.SuccCode)
}

//获取归档数据
func GetArticleRecords(c *gin.Context) {
	var books []models.ArticleRecordNum
	config.Db.Raw("SELECT a.id, a.title, a.create_date,(SELECT count(1) FROM work_review where a_id = a.id) as reviewCount FROM work_article a order by a.create_date desc").Find(&books)
	var rs []models.ArticleRecord
	for _, book := range books {
		var record models.ArticleRecord
		record.Id = book.Id
		record.Title = book.Title
		record.ReviewCount = book.ReviewCount
		record.YearM = time.Unix(int64(book.CreateDate), 0).Format("2006-01")
		record.DateM = time.Unix(int64(book.CreateDate), 0).Format("01-02")
		rs = append(rs, record)
	}
	c.JSON(http.StatusOK, models.RetunMsgFunc(models.SuccCode, rs))
}

//获取留言数据
func GetGuestBooks(c *gin.Context) {
	id, _ := strconv.Atoi(c.Query("id"))
	currentPage, _ := strconv.Atoi(c.Query("currentPage"))
	if currentPage == config.Common_ZERO {
		currentPage = config.Common_ONE
	}
	//留言默认每页请求为5条数据
	returnMap := make(map[string]interface{}, config.Common_TWO)
	var books []models.GuestBook
	config.Db.Table("work_review").Order("create_time desc").Where("a_id = ?", id).Limit(config.Common_FIVE).Offset((currentPage - config.Common_ONE) * config.Common_FIVE).Find(&books)
	returnMap["list"] = books
	var maxSize int
	config.Db.Raw("select count(1) from work_review where id = ?", id).Scan(&maxSize)
	maxPage := 0
	if maxSize%config.Common_FIVE == config.Common_ZERO {
		maxPage = maxSize / config.Common_FIVE
	} else {
		maxPage = maxSize/config.Common_FIVE + config.Common_ONE
	}
	returnMap["maxPage"] = maxPage
	c.JSON(http.StatusOK, models.RetunMsgFunc(models.SuccCode, books))
}
