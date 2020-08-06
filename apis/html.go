package apis

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func IndexHtml(c *gin.Context) {
	c.HTML(http.StatusOK, "index.html", "")
}

func GuestBookHtml(c *gin.Context) {
	c.HTML(http.StatusOK, "gustbook.html", "")
}

func ArticleHtml(c *gin.Context) {
	c.HTML(http.StatusOK, "fuwenben.html", "")
}

func SearchHtml(c *gin.Context) {
	c.HTML(http.StatusOK, "search.html", "")
}

func ArchivesHtml(c *gin.Context) {
	c.HTML(http.StatusOK, "archives.html", "")
}

func DetailHtml(c *gin.Context) {
	c.HTML(http.StatusOK, "detail.html", "")
}

func LinkHtml(c *gin.Context) {
	c.HTML(http.StatusOK, "link.html", "")
}

func UpdateHtml(c *gin.Context) {
	c.HTML(http.StatusOK, "update.html", "")
}
