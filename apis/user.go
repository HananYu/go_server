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

func ChangePassword(c *gin.Context) {
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
	user.Salt = utils.GetRandomString(6)
	user.Password = strings.ToLower(utils.MD5([]byte(user.Password + user.Salt)))
	config.Db.Table("sys_user").Where("account=?", user.Account).Update(map[string]interface{}{"password": user.Password, "salt": user.Salt})
	c.JSON(http.StatusOK, models.SuccCode)
}

//注册用户
func InsertUser(c *gin.Context) {
	var user models.User
	err := c.BindJSON(&user)
	if err != nil {
		c.JSON(http.StatusOK, models.ReqCode)
		return
	}
	if user.Account == "" || user.Password == "" {
		c.JSON(http.StatusBadRequest, models.ReqCode) //作为错误处理
		return
	}
	config.Db.Table("sys_user").Where("account = ?", user.Account).First(&user)
	if user.Id != config.CommonZero {
		//账号已经存在，不能进行注册
		c.JSON(http.StatusOK, models.AccountCode)
		return
	}
	user.CreateDate = int(time.Now().Unix()) //获取时间戳
	user.Salt = utils.GetRandomString(6)
	user.Password = strings.ToLower(utils.MD5([]byte(user.Password + user.Salt)))

	config.Db.Table("sys_user").Create(&user)
	c.JSON(http.StatusOK, models.SuccCode)
}

//判断用户名是否存在
func FindByAccount(c *gin.Context) {
	account := c.Query("account")
	var user models.User
	config.Db.Table("sys_user").Where("account = ?", account).First(&user)
	if user.Id != config.CommonZero {
		//账号已经存在，不能进行注册
		c.JSON(http.StatusOK, models.AccountCode)
		return
	}
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
	if strings.ToLower(utils.MD5([]byte(user.Password+ur.Salt))) == strings.ToLower(ur.Password) {
		//表示密码一致，则通过，返回token和用户ID
		token := utils.CreaToken(ur.Id, ur.Salt)
		var ut models.UserToken
		ut.Token = token
		ut.UserId = ur.Id
		ut.CreateDate = int(time.Now().Unix())      //10位时间戳精确到秒
		ut.EndDate = ut.CreateDate + (12 * 60 * 60) //有效时间为12小时
		config.Db.Table("sys_user_token").Create(&ut)
		c.JSON(http.StatusOK, models.RetunMsgFunc(models.SuccCode, token))
		return
	}
	c.JSON(http.StatusOK, models.LoginCode)
}

func GetHttp(c *gin.Context) {
	//如果进入表示用户数据校验正确，直接返回通过
	c.JSON(http.StatusOK, models.SuccCode)
}
