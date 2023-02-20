package main

import (
	"fmt"
	"go-tcp/gnet"
	"io"
	"net"
	"time"
)

func main() {
	fmt.Println("client start")
	conn, err := net.Dial("tcp", "127.0.0.1:8888")
	if err != nil {
		fmt.Println("dial error", err)
		return
	}
	for {
		dp := gnet.NewDataPackage()
		binaryMsg, err := dp.Pack(gnet.NewMessage(1, []byte("hi\n")))
		if err != nil {
			fmt.Println("packate false")
			return
		}
		if _, err := conn.Write(binaryMsg); err != nil {
			fmt.Println("write error", err)
			return
		}
		time.Sleep(time.Second * 1)
		fmt.Println("first sleep")
		binaryHead := make([]byte, dp.GetHeadLen())
		if _, err := io.ReadFull(conn, binaryHead); err != nil {
			fmt.Println("read head error", err)
			break
		}
		fmt.Println("readfull finish")
		msgHead, err := dp.Unpack(binaryHead)
		if err != nil {
			fmt.Println("unpack error", err)
		}
		if msgHead.GetMsgLen() > 0 {
			msg := msgHead.(*gnet.Message)
			msg.Data = make([]byte, msgHead.GetMsgLen())
			io.ReadFull(conn, msg.Data)
			fmt.Println("response is ", msg.Data)
		}
		fmt.Println("one round")
		time.Sleep(time.Second * 1)

	}

}
