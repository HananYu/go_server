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
