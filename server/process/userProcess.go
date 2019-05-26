package process

import (
	"chat/common/message"
	"chat/common/utils"
	"encoding/json"
	"fmt"
	"net"
)

type UserProcess struct {
	Conn net.Conn
}

//serverProcessLogin 处理登陆业务
func (u *UserProcess)ServerProcessLogin (mes *message.Message)(err error){

	//获取用户登陆信息
	var loginMes message.LoginMessage
	err = json.Unmarshal([]byte(mes.Data),&loginMes)
	if err != nil{
		fmt.Errorf("json.Unmarshal失败:%v",err)
		return
	}

	//1 声明一个resMes
	var resMes message.Message
	resMes.Type = message.LoginResMesType

	//2 声明一个 LoginResMes

	var LoginResMes message.LoginResMessage

	//假设某一用户ID为100 密码为123456
	if loginMes.UserId==100 && loginMes.UserPwd=="123456" {
		LoginResMes.Code=200
		LoginResMes.Error=""
	} else {
		LoginResMes.Code=404  //用户不存在
		LoginResMes.Error="该用户不存在"
	}
	//3 序列化 LoginResMes 压入 loginMes
	data,err := json.Marshal(LoginResMes)
	if err !=nil{
		fmt.Errorf("序列化 LoginResMes失败:%v",err)
		return
	}
	resMes.Data=string(data)

	//4 序列化 resMes

	data,err = json.Marshal(resMes)
	if err !=nil{
		fmt.Errorf("序列化 resMes失败:%v",err)
		return
	}

	//5 发resMes给客户端
	TransferUtils := &utils.TransferUtils{
		Conn: u.Conn,
	}
	err= TransferUtils.WriterData(data)
	if err !=nil{
		fmt.Errorf("conn.Write失败:%v",err)
		return
	}

	return
}