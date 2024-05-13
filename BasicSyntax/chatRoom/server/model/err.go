package model

import "errors"

// 根据业务需求自定义一些错误。。。
var (
	ERROR_USER_NOEEXIST = errors.New("用户不存在。。。")
	ERROR_USER_EXIST    = errors.New("用户已存在。。。")
	ERROR_USER_PWD      = errors.New("密码错误。。。")
)
