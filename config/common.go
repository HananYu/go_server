package config

const (
	//配置文件----开发
	SavePathUrl string = "F:/Temp/"                                                              //文件保存路径
	ServiceUrl  string = "http://192.168.99.198:8000/file/"                                      //文件访问地址
	MysqlUrl    string = "root:921115@tcp(192.168.99.199:3306)/yusj?charset=utf8&parseTime=true" //数据库配置

	//配置文件----上线
	//SavePathUrl string = "/home/temp/"                                                   //文件保存路径
	//ServiceUrl  string = "https://www.hananyu.cn/file/"                                  //文件访问地址
	//MysqlUrl string = "root:921115@tcp(127.0.0.1:3306)/yusj?charset=utf8&parseTime=true" //数据库配置

	//token常量
	TokenUSERID string = "userId"
	TokenSALT   string = "salt"
	TokenTOKEN  string = "token"

	//表名
	DataTableUser    string = "sys_user"
	DataTableMenu    string = "sys_menu"
	DataTableCode    string = "sys_user_code"
	DataTableToken   string = "sys_user_token"
	DataTableArticle string = "work_article"
	DataTableReview  string = "work_review"

	//字符串常量
	CommonNull            string = ""
	CommonId              string = "id"
	CommonName            string = "name"
	CommonAccount         string = "account"
	CommonPassword        string = "password"
	CommonFile            string = "file"
	CommonList            string = "list"
	CommonArticleTypeZero string = "随笔"
	CommonArticleTypeOne  string = "随记"
	CommonArticleTypeTwo  string = "随行"
	CommonCurrentPage     string = "currentPage"
	CommonPageSize        string = "pageSize"
	CommonImageIndex      string = "<img src="
	CommonImageLast       string = ">"

	//符号
	CommonComma string = ","

	//图片常量
	CommonPng string = "png"
	CommonGif string = "gif"

	//bool
	CommonTrue  bool = true
	CommonFalse bool = false

	//常量数字
	CommonZero  int = 0
	CommonOne   int = 1
	CommonTwo   int = 2
	CommonThree int = 3
	CommonFour  int = 4
	CommonFive  int = 5
)
