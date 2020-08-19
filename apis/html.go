package apis

import (
	"gin/config"
	"github.com/gin-gonic/gin"
	"net/http"
)

func LoginHtml(c *gin.Context) {
	c.HTML(http.StatusOK, "login.html", config.CommonNull)
}

func IndexHtml(c *gin.Context) {
	c.HTML(http.StatusOK, "index.html", config.CommonNull)
}

func GuestBookHtml(c *gin.Context) {
	c.HTML(http.StatusOK, "gustbook.html", config.CommonNull)
}

func ArticleHtml(c *gin.Context) {
	c.HTML(http.StatusOK, "fuwenben.html", config.CommonNull)
}

func SearchHtml(c *gin.Context) {
	c.HTML(http.StatusOK, "search.html", config.CommonNull)
}

func ArchivesHtml(c *gin.Context) {
	c.HTML(http.StatusOK, "archives.html", config.CommonNull)
}

func DetailHtml(c *gin.Context) {
	c.HTML(http.StatusOK, "detail.html", config.CommonNull)
}

func LinkHtml(c *gin.Context) {
	c.HTML(http.StatusOK, "link.html", config.CommonNull)
}

func UpdateHtml(c *gin.Context) {
	c.HTML(http.StatusOK, "update.html", config.CommonNull)
}

func ModifyHtml(c *gin.Context) {
	c.HTML(http.StatusOK, "modify.html", config.CommonNull)
}
