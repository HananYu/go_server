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
