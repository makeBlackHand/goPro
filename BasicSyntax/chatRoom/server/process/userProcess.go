package process

import (
	"encoding/json"
	"fmt"
	"goPro/BasicSyntax/chatRoom/common/message"
	"goPro/BasicSyntax/chatRoom/server/model"
	"goPro/BasicSyntax/chatRoom/server/util"
	"log"
	"net"
)

type UserProcess struct {
	Con    net.Conn
	UserId int //表示是哪个用户的con
}

// 通知所有在线用户方法
func (this *UserProcess) NotifyOtherOnlineUser(UserId int) {
	//遍历Mgr中的users切片一个个发送NotifyUserStatusMes
	for id, up := range userMgr.onlineUser {
		if id == UserId {
			continue
		}
		up.NotifyMeOnline(UserId)
	}
}

func (this *UserProcess) NotifyMeOnline(UserId int) {
	var mes message.Message
	mes.Type = message.NotifyUserStatusMesType
	var notifyUserStatusMes message.NotifyUserStatusMes
	notifyUserStatusMes.UserId = UserId
	notifyUserStatusMes.Status = message.UserOnline
	//将notifyUserStatusMes序列化后放入mes.data
	data, err := json.Marshal(notifyUserStatusMes)
	if err != nil {
		fmt.Println("notifyUserStatusMes fail")
		return
	}
	//转string类型放入mes.data
	mes.Data = string(data)
	//再将mes序列化后准备发送
	data, err = json.Marshal(mes)
	if err != nil {
		fmt.Println("notifyUserStatusMes fail")
		return
	}
	tf := &util.Transfer{
		Con: this.Con,
	}
	err = tf.WritePkg(data)
	if err != nil {
		fmt.Println("notifyMeOnline writePkg fail")
		return
	}

}

func (this *UserProcess) ServerProcessRegister(mes *message.Message) (err error) {
	var registerMes message.RegisterMes                  //把mse.data的数取出来
	err = json.Unmarshal([]byte(mes.Data), &registerMes) //把mes.data反序列化放进registerMes
	if err != nil {
		fmt.Println("serverProcessLogin 反序列化失败", err)
		return
	}
	var resMes message.Message //先声明一个message结构体
	resMes.Type = message.RegisterResMesType

	var registerResMes message.RegisterResMes //再声明一个loginResMes的结构体
	err = model.MyUserDao.Register(&registerMes.User)
	if err != nil {
		if err == model.ERROR_USER_EXIST {
			registerResMes.Code = 505
			registerResMes.Error = model.ERROR_USER_EXIST.Error()
		} else {
			registerResMes.Code = 506
			registerResMes.Error = "注册发生位置错误"
		}
	} else {
		registerResMes.Code = 200
	}
	data, err := json.Marshal(registerResMes)
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
	}
	err = tf.WritePkg(data)
	return
}

// 专门处理登录请求
func (this *UserProcess) ServerProcessLogin(mes *message.Message) (err error) {
	log.Println("ServerProcessLogin")
	var loginMes message.LoginMes //把mse.data的数取出来

	err = json.Unmarshal([]byte(mes.Data), &loginMes) //把mes.data反序列化放进loginMes
	if err != nil {
		fmt.Println("serverProcessLogin 反序列化失败", err)
		return
	}
	var resMes message.Message //先声明一个message结构体
	resMes.Type = message.LoginResMestype

	var loginResMes message.LoginResMes //再声明一个loginResMes的结构体

	user, err := model.MyUserDao.Login(loginMes.UserID, loginMes.UserPwd)
	log.Println("ServerProcessLogin.user =", user)

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
		// pass
		loginResMes.Code = 200

		//将登录成功的id赋值给this
		this.UserId = loginMes.UserID
		//登录成功的放入UserMgr
		userMgr.Add(this)
		//通知其他在线用户，我上线了
		this.NotifyOtherOnlineUser(loginMes.UserID)
		for id, _ := range userMgr.onlineUser {
			loginResMes.UserIds = append(loginResMes.UserIds, id)
		}
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
