package main

import (
	"fmt"
	"goPro/BasicSyntax/chatRoom/server/model"
	"net"
	"time"
)

func IninUserDao() {
	//Pool本身就是全局变量可以直接调用
	//要注意初始化问题 要先调用initPool在调用initUserDao
	model.MyUserDao = model.NewUserDao(Pool)
}

// main负责监听，初始化，等待连接等
func process(con net.Conn) {
	defer con.Close()
	p := &Processor{
		Con: con,
	}
	err := p.Process2()
	if err != nil {
		fmt.Println("服务器和客户端协程错误", err)
		return
	}
}

func Init() {
	//服务器启动时就初始化redis的连接池
	InitPool("api.ikanned.com:16379", 16, 0, 300*time.Second)
	IninUserDao()
}

func main() {
	Init()
	println("服务器在[新的结构]8889端口监听")
	listen, err := net.Listen("tcp", "localhost:8889")
	defer listen.Close()
	if err != nil {
		fmt.Println("listen 错误", err)
		return
	}
	for {
		fmt.Println("等待客户端连接服务器")
		con, error := listen.Accept()
		if error != nil {
			fmt.Println("连接失败")
		}
		go process(con)
	}

}

// readPkg 用于读取数据包
//func ReadPkg2(con net.Conn) (mes message.Message, err error) {
//	buffer := make([]byte, 8096)
//	log.Println("读取客户端发送的数据...")
//	// 这里的conn只有在没有被关闭的情况下 才会阻塞
//	// 如果客户端关闭了conn就不会阻塞了
//	_, err = con.Read(buffer[:4]) // 先读取前面四个字节
//	if err != nil {
//		log.Println("read err:", err) // 这里是读取头部长度出错
//		//err = errors.New("read pkg header err")
//		// 进行错误处理
//		return
//	}
//	log.Println("读取到buf长度 = ", buffer[:4])
//	// 将这个长度转换为uint32
//	var pkgLen uint32
//	pkgLen = binary.BigEndian.Uint32(buffer[0:4]) // 这里是Uint32转回来 前面是Putuint32
//	// 根据pkgLen来读取内通
//	n, err := con.Read(buffer[:pkgLen]) // 从conn套接字中读取[:pkgLen]长度的内容到buffer里
//	log.Println("buf = ", string(buffer))
//	fmt.Println("n=", n, "pkglen=", pkgLen)
//	if n != int(pkgLen) || err != nil { // 第一种可能读到了但是没有读全 还有一种就是对方关闭了管道
//		log.Println("read err:", err) // 这里是读取内容出错
//		//err = errors.New("read pkg body err")
//		return
//	} // 到这里buffer里面就有收到的消息了 但是还没完 这需要反序列化
//	err = json.Unmarshal(buffer[:pkgLen], &mes) // 注意这里要加上& 很重要 很重要 很重要
//	if err != nil {
//		log.Println("反序列化失败")
//		return
//	}
//	return
//}

//func ReadPkg(con net.Conn) (mes message.Message, error error) {
//	buf := make([]byte, 1024)
//	_, error = con.Read(buf[:4]) //0-3存放消息长度
//	if error != nil || error == io.EOF {
//		//error = errors.New("read pkgHead err")
//		return
//	}
//	var pkgLen uint32
//	pkgLen = binary.BigEndian.Uint32(buf[:4]) //把0-3切片内容长度转换成uint
//	n, error := con.Read(buf[:pkgLen])
//	fmt.Println("n=", n, "pkglen=", pkgLen)
//	if error != nil { //n != int(pkgLen) ||
//		//error = errors.New("read pkgbody err")
//		return
//	}
//	fmt.Println("buf=", string(buf))
//	errors := json.Unmarshal(buf[:pkgLen], &mes)
//	if errors != nil {
//		fmt.Println("反序列化失败", errors)
//		return
//	}
//	return
//	//fmt.Println("读到了", buf[:n])
//}
//
//func writePkg(con net.Conn, data []byte) (err error) {
//	var pkgLen uint32
//	pkgLen = uint32(len(data))
//	log.Println("pkgLen =", pkgLen)
//	var buf [4]byte
//	binary.BigEndian.PutUint32(buf[0:4], pkgLen) //转成MesData字节长的切片
//	log.Println("buf[0.4] = ", binary.BigEndian.Uint32(buf[0:4]))
//	n, e := con.Write(buf[:4])
//	//log.Println("n =", n)
//	if n != 4 || e != nil {
//		fmt.Println("con write err", e)
//		return
//	}
//	n, e = con.Write(data)
//	//log.Println("n =", n)
//	if n != int(pkgLen) || e != nil {
//		fmt.Println("con write err", e)
//		return
//	}
//	return
//}

//// 专门处理登录请求
//func serverProcessLogin(con net.Conn, mes *message.Message) (err error) {
//	var loginMes message.LoginMes                     //把mse.data的数取出来
//	err = json.Unmarshal([]byte(mes.Data), &loginMes) //把mes.data反序列化放进loginMes
//	if err != nil {
//		fmt.Println("serverProcessLogin 反序列化失败", err)
//		return
//	}
//	var resMes message.Message //先声明一个message结构体
//	resMes.Type = message.LoginResMestype
//
//	var loginResMes message.LoginResMes //再声明一个loginResMes的结构体
//
//	if loginMes.UserID == 100 && loginMes.UserPwd == "123abc" {
//		loginResMes.Code = 200 //200合法
//		println("合法")
//	} else {
//		loginResMes.Code = 500
//		println("不合法") //500不合法
//		loginResMes.Error = "用户不存在，请注册后使用！"
//	}
//
//	data, err := json.Marshal(loginResMes)
//	if err != nil {
//		fmt.Println("server-loginResMes-marshal err", err)
//		return
//	}
//	resMes.Data = string(data)
//	data, err = json.Marshal(resMes)
//	if err != nil {
//		fmt.Println("server-ResMes-marshal err", err)
//		return
//	}
//	err = writePkg(con, data)
//	return
//}

//// 根据传的数据不同，决定调用哪个函数
//func serverProcessMes(con net.Conn, mes *message.Message) (err error) {
//	//根据客户端发送的消息种类不同来处理
//	switch mes.Type {
//	case message.LoginMesType: //处理登录逻辑
//		err = serverProcessLogin(con, mes)
//	case message.ResigterMesType: //处理注册逻辑
//	default:
//		fmt.Println("消息类型存在，无法处理。。")
//	}
//	return
//}
