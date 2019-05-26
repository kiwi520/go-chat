package utils

import (
	"chat/common/message"
	"encoding/binary"
	"encoding/json"
	"errors"
	"fmt"
	"net"
)

type TransferUtils struct {
	Conn net.Conn
	Buf [8096]byte
}

//readData 读取数据
func (ut *TransferUtils)ReadData()(mes message.Message,err error){
	//buf := make([]byte,8096)
	_,err=ut.Conn.Read(ut.Buf[:4])

	if  err != nil {
		//err = errors.New("获取发送数据长度失败:"+err.Error())
		return
	}
	var pkglen uint32
	pkglen = binary.BigEndian.Uint32(ut.Buf[0:4])
	//获取客户端发送的数据
	n,err:=ut.Conn.Read(ut.Buf[:pkglen])

	if uint32(n) != pkglen || err != nil{
		fmt.Println("conn.Read err=",err)
		return
	}

	fmt.Println("读到的buf=",string(ut.Buf[:pkglen]))

	//反序列化数据
	err= json.Unmarshal(ut.Buf[:pkglen],&mes)

	if err != nil{
		err = errors.New("获取发送数据失败:"+err.Error())
		return
	}
	return
}


//writerData 给客户端返回数据
func (ut *TransferUtils)WriterData(data []byte)(err error)  {
	// 发送长度给客户端
	var pkglen uint32 = uint32(len(data))
	binary.BigEndian.PutUint32(ut.Buf[0:4],pkglen)
	//发送长度
	n,err:=ut.Conn.Write(ut.Buf[:4])

	if n != 4 || err != nil{
		fmt.Println("conn.Write err=",err)
		return
	}

	//发送data本身
	n,err = ut.Conn.Write(data)

	if uint32(n) != pkglen || err != nil{
		fmt.Println("conn.Write err=",err)
		return
	}

	return
}
