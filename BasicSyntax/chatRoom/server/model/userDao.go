package model

import (
	"encoding/json"
	"fmt"
	"github.com/gomodule/redigo/redis"
)

var MyUserDao *UserDao

// 完成对User结构体的各种操作
type UserDao struct {
	Pool *redis.Pool
}

func NewUserDao(pool *redis.Pool) (userdao *UserDao) {
	userdao = &UserDao{
		Pool: pool,
	}
	return
}

func (this *UserDao) GetUserById(con redis.Conn, id int) (user *User, err error) {
	//fmt.Println("id=", id)
	res, err := redis.String(con.Do("HGet", "users", id))
	//fmt.Println("66666666err", err)
	if err != nil {
		if err == redis.ErrNil {
			err = ERROR_USER_NOEEXIST
		}
		return
	}
	//fmt.Printf("555555555user=", res)
	//s := `{"userID":123,"userPwd":"456","userName":""}`
	err = json.Unmarshal([]byte(res), &user)

	if err != nil {
		fmt.Println("userDao反序列化失败", err)
		return
	}
	return
}

func (this *UserDao) Login(userID int, userPwd string) (user *User, err error) {
	con := this.Pool.Get()
	defer con.Close()
	user, err = this.GetUserById(con, userID)
	if err != nil {
		return
	}
	//fmt.Println("4444444user=", user).'
	if user.UserPwd != userPwd {
		err = ERROR_USER_PWD
		return
	}
	return
}
