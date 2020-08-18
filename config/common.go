package config

const (
	//系统设置 开发
	//SavePathUrl string = "F:/Temp/"                         //文件保存路径
	//ServiceUrl  string = "http://192.168.99.198:8000/file/" //文件访问地址

	//系统设置 上线
	SavePathUrl string = "/home/temp/"                  //文件保存路径
	ServiceUrl  string = "https://www.hananyu.cn/file/" //文件访问地址

	//token常量
	TokenUSERID string = "userId"
	TokenSALT   string = "salt"
	TokenTOKEN  string = "token"

	//常量数字
	CommonZero  int = 0
	CommonOne   int = 1
	CommonTwo   int = 2
	CommonThree int = 3
	CommonFour  int = 4
	CommonFive  int = 5
)
