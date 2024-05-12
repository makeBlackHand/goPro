package main

import "github.com/gomodule/redigo/redis"

var pool *redis.Pool

func init() {
	pool = &redis.Pool{
		MaxIdle:     8,
		MaxActive:   0,
		IdleTimeout: 100,
		Dial: func() (redis.Conn, error) {
			return redis.Dial("tcp", "api.ikanned.com:16379")
		},
	}
}
func main() {
	con := pool.Get()
	defer con.Close()
	con.Do("set", "name", "tom猫")
	r, _ := redis.String(con.Do("get", "name"))
	println(r)
	//pool.Close()
	//con1 := pool.Get()
	//defer con1.Close()
	//con.Do("set", "name1", "tom猫1")
	//r1, _ := redis.String(con.Do("get", "name1"))
	//println(r1)
}
