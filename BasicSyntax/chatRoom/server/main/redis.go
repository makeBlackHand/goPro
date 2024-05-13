package main

import (
	"github.com/gomodule/redigo/redis"
	"time"
)

var Pool *redis.Pool

func InitPool(address string, maxIdle, maxActive int, idleTimeOut time.Duration) {
	Pool = &redis.Pool{
		MaxIdle:     maxIdle,
		MaxActive:   maxActive,
		IdleTimeout: idleTimeOut,
		Dial: func() (redis.Conn, error) {
			return redis.Dial("tcp", address)
		},
	}
}
