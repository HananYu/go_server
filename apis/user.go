package apis

import (
	"gin/config"
	"gin/models"
	"gin/utils"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
	"strings"
	"time"
)

//注册用户
func InsertUser(c *gin.Context) {
	var user models.User
	err := c.BindJSON(&user)
	if err != nil {
		c.JSON(http.StatusOK, models.ReqCode)
		return
	}
	if user.Account == "" || user.Password == "" || user.Phone == "" {
		c.JSON(http.StatusBadRequest, models.ReqCode) //作为错误处理
		return
	}
	user.CreateDate = int(time.Now().Unix()) //获取时间戳
	user.Salt = utils.GetRandomString(6)
	user.IsDel = 0
	user.CreateBy = 0

	config.Db.Table("sys_user").Create(&user)
	c.JSON(http.StatusOK, models.SuccCode)
}

//查找用户
func FindUser(c *gin.Context) {
	cid := c.Query("id")
	id, _ := strconv.Atoi(cid)
	var user models.User
	config.Db.Raw("select * from sys_user where id =?", id).Scan(&user)
	if user.Id == 0 {
		c.JSON(http.StatusBadRequest, models.SysCode) //作为错误处理
		return
	}
	c.JSON(http.StatusOK, models.RetunMsgFunc(models.SuccCode, user))
}

//用户登陆
func Login(c *gin.Context) {
	var user models.User
	err := c.BindJSON(&user)
	if err != nil {
		c.JSON(http.StatusOK, models.ReqCode)
		return
	}
	if user.Account == "" || user.Password == "" {
		c.JSON(http.StatusOK, models.ReqCode)
		return
	}
	var ur models.User
	//Table("sys_user").Where("account = ?", user.Account).Select("xxxx).Row().Scan(&xxx)
	config.Db.Table("sys_user").Where("account = ?", user.Account).First(&ur)
	if strings.ToLower(utils.MD5([]byte(user.Password+user.Salt))) == strings.ToLower(ur.Password) {
		//表示密码一致，则通过，返回token和用户ID
		token := utils.CreaToken(ur.Id, ur.Salt)
		//utils.InserRedis(token, ur.Id) /将用户的信息保存到redis中
		scoreMap := make(map[string]interface{}, 3) //3表示map的容量，可不填
		scoreMap["token"] = token
		scoreMap["userId"] = ur.Id
		scoreMap["salt"] = ur.Salt
		c.JSON(http.StatusOK, models.RetunMsgFunc(models.SuccCode, scoreMap))
		return
	}
	c.JSON(http.StatusOK, models.LoginCode)
}

func GetHttp(c *gin.Context) {
	//如果进入表示用户数据校验正确，直接返回通过
	c.JSON(http.StatusOK, models.SuccCode)
}
