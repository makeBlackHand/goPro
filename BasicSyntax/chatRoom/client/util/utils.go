package util

import (
	"encoding/binary"
	"encoding/json"
	"fmt"
	"goPro/BasicSyntax/chatRoom/common/message"
	"io"
	"log"
	"net"
)

type Transfer struct {
	Con net.Conn
	Buf [8096]byte
}

func (this *Transfer) ReadPkg() (mes message.Message, error error) {
	//buf := make([]byte, 1024)
	_, error = this.Con.Read(this.Buf[:4]) //0-3存放消息长度
	if error != nil || error == io.EOF {
		//error = errors.New("read pkgHead err")
		return
	}
	var pkgLen uint32
	pkgLen = binary.BigEndian.Uint32(this.Buf[:4]) //把0-3切片内容长度转换成uint
	n, error := this.Con.Read(this.Buf[:pkgLen])
	fmt.Println("n=", n, "pkglen=", pkgLen)
	if error != nil { //n != int(pkgLen) ||
		//error = errors.New("read pkgbody err")
		return
	}
	//fmt.Println("buf=", string(this.buf))
	errors := json.Unmarshal(this.Buf[:pkgLen], &mes)
	if errors != nil {
		fmt.Println("反序列化失败", errors)
		return
	}
	return
	//fmt.Println("读到了", buf[:n])
}

func (this *Transfer) WritePkg(data []byte) (err error) {
	var pkgLen uint32
	pkgLen = uint32(len(data))
	log.Println("pkgLen =", pkgLen)
	//var buf [4]byte
	binary.BigEndian.PutUint32(this.Buf[0:4], pkgLen) //转成MesData字节长的切片
	log.Println("buf[0.4] = ", binary.BigEndian.Uint32(this.Buf[0:4]))
	n, e := this.Con.Write(this.Buf[:4])
	//log.Println("n =", n)
	if n != 4 || e != nil {
		fmt.Println("con write err", e)
		return
	}
	n, e = this.Con.Write(data)
	//log.Println("n =", n)
	if n != int(pkgLen) || e != nil {
		fmt.Println("con write err", e)
		return
	}
	return
}
