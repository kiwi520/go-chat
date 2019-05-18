package main

import (
	"encoding/binary"
	"encoding/json"
	"fmt"
	"net"
	"chat/common/message"
)

func Login(userId int,userPwd string) (err error)  {
	//fmt.Printf("userId =%d userPwd =%s",userId,userPwd)
	//return nil

	//链接服务器
	conn,err:= net.Dial("tcp","169.254.223.112:8789")

	if err != nil{
		fmt.Println("net.Dial err=",err)
		return
	}

	//延时关闭
	defer conn.Close()

	//准备通过conn发送消息给服务
	var mes message.Message
	mes.Type = message.LoginMesType
	//创建LoginMes 结构体
	var loginMes message.LoginMessage
	loginMes.UserId = userId
	loginMes.UserPwd = userPwd
	//创建LoginMes 并序列化
	data,err := json.Marshal(loginMes)

	if err != nil{
		fmt.Println("json.Marshal err=",err)
		return
	}

	fmt.Errorf("data内容:%v",data)
	fmt.Errorf("data长度为:%v",len(data))

	//把data数据序列化
	mes.Data =string(data)

	data,err = json.Marshal(mes)

	if err != nil{
		fmt.Println("json.Marshal err=",err)
		return
	}

	//计算data的长度 并发送给服务器
    var pkgLen uint32
	pkgLen = uint32(len(data))
	var buf [4]byte
	binary.BigEndian.PutUint32(buf[0:4],pkgLen)

	//发送长度
	n,err:=conn.Write(buf[:4])
    if n !=4 || err != nil{
    	fmt.Println("conn.Write(bytes) fail",err)
    	return
	}

	fmt.Println("客户端，发送的长度成功")

	//发送长度
	_,err=conn.Write(data)
	if  err != nil{
		fmt.Println("conn.Write(bytes) fail",err)
		return
	}
	fmt.Println("客户端，发送的消息成功")
	return
}
