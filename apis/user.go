package apis

import (
	"fmt"
	"gin/config"
	"gin/models"
	"gin/utils"
	"github.com/gin-gonic/gin"
	"io/ioutil"
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
	user.CreateDate = int(time.Now().Unix() / 1000) //获取时间戳
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
	config.Db.Table("sys_user").Where("account = ?", user.Account).Select("password").First(&ur)
	if strings.ToLower(utils.MD5([]byte(user.Password+user.Salt))) == strings.ToLower(ur.Password) {
		//表示密码一致，则通过，返回token和用户ID
		token := utils.CreaToken(ur.Id, user.Salt)
		utils.InserRedis(token, ur.Id)
		c.JSON(http.StatusOK, models.RetunMsgFunc(models.SuccCode, token))
		return
	}
	c.JSON(http.StatusOK, models.LoginCode)
}


func GetHttp(c *gin.Context) {
	rsp, err := http.Get("http://www.wh.ccoo.cn/tieba/today-9-1-1.html")
	if err != nil {
		// 处理异常
		fmt.Println("请求错误")
		return
	}
	defer rsp.Body.Close()
	body, err := ioutil.ReadAll(rsp.Body) // 读取Body
	//c.HTML(http.StatusOK, string(body), "")
	//fmt.Println(string(body))
	//html,_ := utils.GbkToUtf8(body)
	c.JSON(http.StatusOK, string(body))
}
