package main

import (
	"chat/common/message"
	"chat/common/utils"
	pr "chat/server/process"
	"fmt"
	"io"
	"net"
)

type Processor struct {
	Conn net.Conn
}

//serverProcessMes 根据客户端发送的消息类型调用相应的处理函数
func (p * Processor)serverProcessMes(mes *message.Message)(err error){
	switch mes.Type {
	case message.LoginMesType:
		//登陆处理
		var pr =pr.UserProcess{p.Conn}
		err = pr.ServerProcessLogin(mes)
	case message.RegisterMesType:
		//注册处理
	default:
		fmt.Println("消息类型不存在，无法处理。。。")
	}
	return
}

func (p *Processor)distribution()  {
	//循环读取客户端发送的数据

	for {
		var rd = utils.TransferUtils{
			Conn: p.Conn,
		}
		data,err:=rd.ReadData()
		if  err != nil{

			if err == io.EOF {
				fmt.Println("客户端退出，服务器也正常关闭")
				return
			}else{
				fmt.Println("conn.Read err=",err)
				return
			}
		}

		err = p.serverProcessMes(&data)

		if err != nil{
			//fmt.Println("conn.Read err=",err)
			return
		}

		fmt.Println("读到的buf=",data.Data)
	}
}