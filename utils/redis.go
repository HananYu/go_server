package utils

import (
	"fmt"
	"github.com/garyburd/redigo/redis"
)

func InitRedis() (redis.Conn, error) {
	c, err := redis.Dial("tcp", "127.0.0.1:6379")

	return c, err
	//defer c.Close()
}

//插入数据
func InserRedis(mkey string, mvalue interface{}) bool {
	c, err := InitRedis()
	defer c.Close()
	if err != nil {
		fmt.Println("Connect to redis error", err)
		return false
	}
	c.Do("SET", mkey, mvalue)
	return true
}

//查找数据
func FindRedis(mkey string) string {
	c, err := InitRedis()
	defer c.Close()
	if err != nil {
		fmt.Println("Connect to redis error", err)
		return ""
	}
	value, _ := redis.String(c.Do("GET", mkey))
	return value
}

//判断redis中是否存在
func ExisitRedis(mkey string) bool {
	c, err := InitRedis()
	defer c.Close()
	if err != nil {
		fmt.Println("Connect to redis error", err)
		return false
	}
	value, _ := redis.Bool(c.Do("EXISTS", mkey))
	return value
}