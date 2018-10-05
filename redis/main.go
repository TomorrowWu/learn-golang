package redis

import (
	"fmt"

	"github.com/garyburd/redigo/redis"
)

func main() {
	//通过go向redis写入数据和读取数据
	//1,链接到redis
	conn, err := redis.Dial("tcp", "127.0.0.1:6379")
	if err != nil {
		fmt.Println("redis.Dial err=", err)
		return
	}
	defer conn.Close()

	_, err = conn.Do("Set", "name", "tomjerry毛毛")
	if err != nil {
		fmt.Println("redis.Set err=", err)
		return
	}

	res, err := redis.String(conn.Do("Get", "name"))
	if err != nil {
		fmt.Println("redis.Get err=", err)
		return
	}

	fmt.Println("操作OK,res=", res)

}
