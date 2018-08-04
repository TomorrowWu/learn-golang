package main

import (
	"time"

	"github.com/garyburd/redigo/redis"
)

var pool *redis.Pool

func initPool(address string, maxIdle, maxActive int, idleTimeout time.Duration) {
	pool = &redis.Pool{
		Dial: func() (redis.Conn, error) {
			return redis.Dial("tcp", address)
		},
		//TestOnBorrow: nil,
		MaxIdle:     maxIdle,     //最大空闲连接数
		MaxActive:   maxActive,   //表示和数据库的最大连接数 0表示没有限制
		IdleTimeout: idleTimeout, //最大空闲时间
		//Wait:        false,
		//MaxConnLifetime: 0,
	}
}
