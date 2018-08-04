package model

import (
	"encoding/json"

	"fmt"

	"github.com/garyburd/redigo/redis"
)

//服务器启动后,就初始化一个UserDao实例
//全局变量,在需要和redis操作时,直接使用
var (
	MyUserDao *UserDao
)

type UserDao struct {
	pool *redis.Pool
}

//使用工厂模式,创建一个UserDao实例
func NewUserDao(pool *redis.Pool) (userDao *UserDao) {
	userDao = &UserDao{
		pool: pool,
	}
	return
}

func (userDao *UserDao) getUserById(conn redis.Conn, id int) (user *User, err error) {
	res, err := redis.String(conn.Do("HGet", "users", id))
	if err != nil {
		if err == redis.ErrNil {
			err = ERROR_USER_NOTEXITS
		}
		return
	}

	user = &User{}

	err = json.Unmarshal([]byte(res), user)
	if err != nil {
		fmt.Println("json.Unmarshal err=", err)
		return
	}

	return
}

//Login 完成登录的校验
func (userDao *UserDao) Login(userId int, userPwd string) (user *User, err error) {
	conn := userDao.pool.Get()
	defer conn.Close()

	user, err = userDao.getUserById(conn, userId)
	if err != nil {
		return
	}

	if user.UserPwd != userPwd {
		err = ERROR_USER_PWD
		return
	}
	return
}
