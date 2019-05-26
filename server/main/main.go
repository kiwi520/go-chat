package main

import (
	"fmt"
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

//处理和客户端通信
func process(conn net.Conn)  {
	//延时 关闭
	defer conn.Close()

	var proesspr = &Processor{
		conn,
	}
	proesspr.distribution()
	////循环读取客户端发送的数据
	//
	//for {
	//	data,err:=utils.ReadData(conn)
	//	if  err != nil{
	//
	//		if err == io.EOF {
	//			fmt.Println("客户端退出，服务器也正常关闭")
	//			return
	//		}else{
	//			fmt.Println("conn.Read err=",err)
	//			return
	//		}
	//	}
	//
	//	err = serverProcessMes(&data)
	//
	//	if err != nil{
	//		//fmt.Println("conn.Read err=",err)
	//		return
	//	}
	//
	//	fmt.Println("读到的buf=",data.Data)
	//}
}

