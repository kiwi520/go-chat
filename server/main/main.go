package main

import (
	"chat/common/message"
	"chat/common/utils"
	"encoding/json"
	"fmt"
	"io"
	"net"
)

func main()  {

	listen,err:= net.Listen("tcp","192.168.0.104:8789")

	if err != nil{
		fmt.Printf("net.Listen err:%v",err)
		return
	}else{
		fmt.Println("8789服务已开启，正在监听中...")
	}


	for {
		fmt.Println("等待客户端链接")
		conn,err := listen.Accept()
		if err != nil {
			fmt.Println("Listen.Accept err =",err)
		}

		//保持与客户端通信

		go process(conn)
	}
}

//serverProcessMes 根据客户端发送的消息类型调用相应的处理函数
func serverProcessMes(conn net.Conn,mes *message.Message)(err error){
	switch mes.Type {
	case message.LoginMesType:
		//登陆处理
		err = serverProcessLogin(conn,mes)
	case message.RegisterMesType:
		//注册处理
	default:
		fmt.Println("消息类型不存在，无法处理。。。")
	}
	return
}

//serverProcessLogin 处理登陆业务
func serverProcessLogin (conn net.Conn,mes *message.Message)(err error){

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
	err= utils.WriterData(conn,[]byte(data))
	if err !=nil{
		fmt.Errorf("conn.Write失败:%v",err)
		return
	}

	return
}

//处理和客户端通信
func process(conn net.Conn)  {
	//延时 关闭
	defer conn.Close()

	//循环读取客户端发送的数据

	for {
		data,err:=utils.ReadData(conn)
		if  err != nil{

			if err == io.EOF {
				fmt.Println("客户端退出，服务器也正常关闭")
				return
			}else{
				fmt.Println("conn.Read err=",err)
				return
			}
		}

		err = serverProcessMes(conn,&data)

		if err != nil{
			//fmt.Println("conn.Read err=",err)
			return
		}

		fmt.Println("读到的buf=",data.Data)
	}
}

