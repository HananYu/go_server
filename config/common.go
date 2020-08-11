package config

const (
	//系统设置 开发
	Save_Path_URL string = "F:/Temp/"                         //文件保存路径
	Service_URL   string = "http://192.168.99.198:8000/file/" //文件访问地址

	//系统设置 上线
	//Save_Path_URL string = "/home/temp/"                         //文件保存路径
	//Service_URL   string = "https://www.hananyu.cn/file/" //文件访问地址

	//token常量
	Token_USERID string = "userId"
	Token_SALT   string = "salt"
	Token_TOKEN  string = "token"

	//常量数字
	Common_ZERO  int = 0
	Common_ONE   int = 1
	Common_TWO   int = 2
	Common_THREE int = 3
	Common_FOUR  int = 4
	Common_FIVE  int = 5
)
