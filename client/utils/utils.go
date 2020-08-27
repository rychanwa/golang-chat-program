package utils

import (
	"net"
	"chat/base"
	"encoding/binary"
	"encoding/json"
	"fmt"
	"io"
)

type Transfer struct {
	Con net.Conn

}

// 发送消息
func (this *Transfer) SendMessage(data []byte) (err error) {
	var messLen uint32
	MessByte := make([]byte, 4)
	messLen = uint32(len(data))
	binary.BigEndian.PutUint32(MessByte[:], messLen)
	n, err := this.Con.Write(MessByte) // 发送消息长度
	if err != nil || n != 4 {
		fmt.Println(err)
		return err
	}
	n, err = this.Con.Write(data) // 发送消息
	if err != nil || n != int(messLen) {
		fmt.Println(err)
		return err
	}
	return
}

// 读取消息内容
func (this *Transfer) ReadMessage() (mess base.Message, err error) {
	var messbyte []byte = make([]byte, 1024)
	_, err = this.Con.Read(messbyte[:4]) //
	if err != nil {
		if err == io.EOF {
			fmt.Println("没有新的消息")
		} else {
			fmt.Println("读取消息长度时，发现错误：", err)
		}
		return
	}
	// fmt.Print(messbyte[:4]) //
	var messLen uint32
	messLen = binary.BigEndian.Uint32(messbyte[:4])
	n, err := this.Con.Read(messbyte[:messLen]) //
	if err != nil || n != int(messLen) {
		fmt.Println("读取消息内容时，发现错误：", err)
		return
	}
	err = json.Unmarshal(messbyte[:messLen], &mess)
	if err != nil {
		fmt.Println("反序列化消息内容时，发现错误：", err)
		return
	}
	return
}
