package main

import (
	"fmt"
	"github.com/gomodule/redigo/redis"
)

func main() {

	conn, _ := redis.Dial("tcp", "api.ikanned.com:16379")
	conn.Do("flushall")
	defer conn.Close()
	conn.Do("set", "name", "zhangsan")
	res, _ := redis.String(conn.Do("get", "name"))
	println("name=", res)
	conn.Do("HSet", "user", "name", "tom")
	res2, _ := redis.String(conn.Do("HGet", "user", "name"))
	println("hash=", res2)
	conn.Do("HSet", "user", "age", 10)
	res3, _ := redis.Int(conn.Do("HGet", "user", "age"))
	println("hash=", res3)

	conn.Do("HMSet", "user", "address", "beijing", "email", "1.com")
	res6, _ := redis.Strings(conn.Do("HMGet", "user", "name", "age", "address", "email"))
	fmt.Println("HMGet=", res6)
	for i, s := range res6 {
		fmt.Printf("r[%v]=%v\n", i, s)
	}
}
