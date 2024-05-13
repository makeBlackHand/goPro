package model

type User struct {
	UserId   int    `json:"userID"` //数据库里面的key是小写的所以要转小写！！！
	UserPwd  string `json:"userPwd"`
	UserName string `json:"userName"`
}
