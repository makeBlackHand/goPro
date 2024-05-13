package process

import (
	"encoding/binary"
	"encoding/json"
	"fmt"
	"goPro/BasicSyntax/chatRoom/client/util"
	"goPro/BasicSyntax/chatRoom/common/message"
	"log"
	"net"
)

type UserProcess struct {
}

func (this *UserProcess) Login(id int, pwd string) (err error) {
	con, err := net.Dial("tcp", "localhost:8889")
	if err != nil {
		fmt.Println("连接失败")
		return
	}
	defer con.Close()
	var mes = message.Message{}
	mes.Type = message.LoginMesType
	var loginMes message.LoginMes
	loginMes.UserID = id
	loginMes.UserPwd = pwd
	data, err := json.Marshal(loginMes)
	if err != nil {
		fmt.Println("序列化失败")
		return
	}
	mes.Data = string(data)
	Mesdata, err := json.Marshal(mes)
	if err != nil {
		fmt.Println("序列化失败")
		return
	}

	var pkgLen uint32
	pkgLen = uint32(len(Mesdata)) //求Mesdata的字节
	log.Println("pkgLen =", pkgLen)
	var buf [4]byte
	binary.BigEndian.PutUint32(buf[0:4], pkgLen) //转成MesData字节长的切片
	log.Println("buf[0.4] = ", binary.BigEndian.Uint32(buf[0:4]))
	n, err := con.Write(buf[:4])
	//log.Println("n =", n)
	if n != 4 || err != nil {
		fmt.Println("con write err", err)
		return
	}
	_, err = con.Write(Mesdata)
	if err != nil {
		fmt.Println("write faill", err)
		return
	}

	var loginResMesData message.LoginResMes
	tf := &util.Transfer{
		Con: con,
		Buf: [8096]byte{},
	}
	mes, err = tf.ReadPkg()
	if err != nil {
		fmt.Println("loginResMes readPkg ERR")
		return
	}
	err = json.Unmarshal([]byte(mes.Data), &loginResMesData)
	log.Println(loginResMesData)
	if loginResMesData.Code == 200 {
		go ServerProcessMes(con)
		for {
			ShowMenu()
		}
		//fmt.Println("登陆成功")
	} else {
		fmt.Println(loginResMesData.Error)
	}
	//time.Sleep(20 * time.Second)
	//fmt.Println("休眠20s")
	fmt.Println("客户端发送消息长度成功,发了", pkgLen, "内容是:", string(Mesdata))
	return
}
