package model

import (
	"encoding/json"
	"fmt"
	"github.com/gomodule/redigo/redis"
	"goPro/BasicSyntax/chatRoom/common/message"
	"log"
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

func (this *UserDao) GetUserById(con redis.Conn, id int) (user *message.User, err error) {
	//fmt.Println("id=", id)
	res, err := redis.String(con.Do("HGet", "users", id))

	fmt.Println("77777777777777user", res)
	if err != nil {
		if err == redis.ErrNil {
			err = ERROR_USER_NOEEXIST
		}
		return
	}
	//fmt.Printf("555555555user=", res)
	//s := `{"userID":123,"userPwd":"456","userName":""}`
	err = json.Unmarshal([]byte(res), &user)
	//fmt.Println("user=", user)
	if err != nil {
		fmt.Println("userDao反序列化失败", err)
		return
	}
	log.Println("GetUserById =", user)
	return
}

func (this *UserDao) Login(userID int, userPwd string) (user *message.User, err error) {
	log.Println("Handle Login")
	con := this.Pool.Get()
	defer con.Close()
	user, err = this.GetUserById(con, userID)
	if err != nil {
		return
	}
	log.Println("Login.user =", user)
	//fmt.Println("4444444user=", user).'
	if user.UserPwd != userPwd {
		err = ERROR_USER_PWD
		log.Println("密码错误")
		return
	}
	return user, nil
}

func (this *UserDao) Register(user *message.User) (err error) {
	con := this.Pool.Get()
	defer con.Close()
	_, err = this.GetUserById(con, user.UserId)
	if err == nil {
		err = ERROR_USER_EXIST //取用户信息没有错误代表用户已存在
		return
	}
	//fmt.Println("4444444user=", user)
	//if user.UserPwd != user.UserPwd {
	//	err = ERROR_USER_PWD
	//	return
	//}
	data, err := json.Marshal(user)
	if err != nil {
		fmt.Println("registerMes 序列化失败")
		return
	}
	_, err = con.Do("HSet", "users", user.UserId, string(data))
	if err != nil {
		fmt.Println("保存用户错误")
		return
	}
	return
}
