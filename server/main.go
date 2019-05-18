package main

import (
	"chat/common/message"
	"encoding/binary"
	"encoding/json"
	"fmt"
	"net"
)

func main()  {

	listen,err:= net.Listen("tcp","169.254.223.112:8789")

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

	//循环读取客户端发送的数据

	for {
		//buf := make([]byte,8096)
		//_,err:=conn.Read(buf[:4])
		//
		//if  err != nil{
		//	fmt.Println("conn.Read err=",err)
		//	return
		//}
		data,err:=readData(conn)
		if  err != nil{
			fmt.Println("conn.Read err=",err)
			return
		}

		fmt.Println("读到的buf=",data.Data)
	}
}

func readData(conn net.Conn)(mes message.Message,err error){
	buf := make([]byte,8096)
	_,err=conn.Read(buf[:4])

	if  err != nil{
		fmt.Println("conn.Read err=",err)
		return
	}
	var pkglen uint32
	pkglen = binary.BigEndian.Uint32(buf[0:4])

	n,err:=conn.Read(buf[:pkglen])

	if uint32(n) != pkglen || err != nil{
		fmt.Println("conn.Read err=",err)
		return
	}

	fmt.Println("读到的buf=",string(buf[:pkglen]))

	err= json.Unmarshal(buf[:pkglen],&mes)

	if err != nil{
		fmt.Println("json.Unmarshal err=",err)
		return
	}
	return
}