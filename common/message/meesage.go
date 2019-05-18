package message

const (
	LoginMesType ="LoginMes"
	LoginResMesType ="LoginResMes"
)


type Message struct {
	Type string //消息类型
	Data string //消息内容
}

//登陆结构体
type LoginMessage struct {
	UserId  int //用户id
	UserPwd  string //用户密码
	UserName string //用户昵称
}

type LoginResMessage struct {
	Code int // 返回状态码
	Error string //返回错误信息
}