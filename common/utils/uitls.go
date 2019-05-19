package utils
import (
	"chat/common/message"
	"encoding/binary"
	"encoding/json"
	"errors"
	"fmt"
	"net"
)

//readData 读取数据
func ReadData(conn net.Conn)(mes message.Message,err error){
	buf := make([]byte,8096)
	_,err=conn.Read(buf[:4])

	if  err != nil {
		//err = errors.New("获取发送数据长度失败:"+err.Error())
		return
	}
	var pkglen uint32
	pkglen = binary.BigEndian.Uint32(buf[0:4])
	//获取客户端发送的数据
	n,err:=conn.Read(buf[:pkglen])

	if uint32(n) != pkglen || err != nil{
		fmt.Println("conn.Read err=",err)
		return
	}

	fmt.Println("读到的buf=",string(buf[:pkglen]))

	//反序列化数据
	err= json.Unmarshal(buf[:pkglen],&mes)

	if err != nil{
		err = errors.New("获取发送数据失败:"+err.Error())
		return
	}
	return
}


//writerData 给客户端返回数据
func WriterData(conn net.Conn,data []byte)(err error)  {
	// 发送长度给客户端
	var buf [4]byte
	var pkglen uint32 = uint32(len(data))
	binary.BigEndian.PutUint32(buf[0:4],pkglen)
	//发送长度
	n,err:=conn.Write(buf[:4])

	if n != 4 || err != nil{
		fmt.Println("conn.Write err=",err)
		return
	}

	//发送data本身
	n,err = conn.Write(data)

	if uint32(n) != pkglen || err != nil{
		fmt.Println("conn.Write err=",err)
		return
	}

	return
}
