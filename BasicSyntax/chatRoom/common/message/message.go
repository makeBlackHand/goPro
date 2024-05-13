package message

const (
	LoginMesType    = "LoginMes"
	LoginResMestype = "LoginResMes"
	ResigterMesType = "ResgisterMes"
)

type Message struct {
	Type string `json:"type"`
	Data string `json:"data"`
}

type LoginMes struct {
	UserID   int    `json:"userID"`
	UserPwd  string `json:"userPwd"`
	UserName string `json:"userName"`
}

type LoginResMes struct {
	Code  int    `json:"code"` //返回状态码 500表示用户未注册 200为登录成功
	Error string `json:"error"`
}

type Register struct {
}
