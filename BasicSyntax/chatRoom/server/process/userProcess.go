package process

import (
	"encoding/json"
	"fmt"
	"goPro/BasicSyntax/chatRoom/common/message"
	"goPro/BasicSyntax/chatRoom/server/model"
	"goPro/BasicSyntax/chatRoom/server/util"
	"net"
)

type UserProcess struct {
	Con net.Conn
}

// 专门处理登录请求
func (this *UserProcess) ServerProcessLogin(mes *message.Message) (err error) {
	var loginMes message.LoginMes                     //把mse.data的数取出来
	err = json.Unmarshal([]byte(mes.Data), &loginMes) //把mes.data反序列化放进loginMes
	if err != nil {
		fmt.Println("serverProcessLogin 反序列化失败", err)
		return
	}
	var resMes message.Message //先声明一个message结构体
	resMes.Type = message.LoginResMestype

	var loginResMes message.LoginResMes //再声明一个loginResMes的结构体

	user, err := model.MyUserDao.Login(loginMes.UserID, loginMes.UserPwd)

	if err != nil {
		if err == model.ERROR_USER_NOEEXIST {
			loginResMes.Code = 500
			loginResMes.Error = err.Error() //err.Error()是取出2err里的错误
		} else if err == model.ERROR_USER_PWD {
			loginResMes.Code = 403
			loginResMes.Error = err.Error()
		} else {
			loginResMes.Code = 505
			loginResMes.Error = "服务器内部错误"
		}
	} else {
		loginResMes.Code = 200
		fmt.Println(user, "登录成功")
	}
	//if loginMes.UserID == 100 && loginMes.UserPwd == "123abc" {
	//	loginResMes.Code = 200 //200合法
	//	println("合法")
	//} else {
	//	loginResMes.Code = 500
	//	println("不合法") //500不合法
	//	loginResMes.Error = "用户不存在，请注册后使用！"
	//}

	data, err := json.Marshal(loginResMes)
	if err != nil {
		fmt.Println("server-loginResMes-marshal err", err)
		return
	}
	resMes.Data = string(data)
	data, err = json.Marshal(resMes)
	if err != nil {
		fmt.Println("server-ResMes-marshal err", err)
		return
	}
	tf := &util.Transfer{
		Con: this.Con,
		Buf: [8096]byte{},
	}
	err = tf.WritePkg(data)
	return
}
