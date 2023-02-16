package main

import (
	"fmt"
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
		_, err := conn.Write([]byte("hello server"))
		if err != nil {
			fmt.Println("write error", err)
			return
		}
		buf := make([]byte, 512)
		cnt, err := conn.Read(buf)
		if err != nil {
			fmt.Println("read buf error")
			return
		}
		fmt.Printf("server call back %s\n", buf[:cnt])
		time.Sleep(time.Second * 1)

	}

}
