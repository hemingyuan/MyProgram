package main

import (
	"bufio"
	"fmt"
	"net"
	"os"
)

func dailServer(addr string) (conn net.Conn, err error) {
	conn, err = net.Dial("tcp", addr)
	return
}

func recvMsg(conn net.Conn) {
	for {
		buf := make([]byte, 1024)
		n, err := conn.Read(buf)
		if err != nil {
			fmt.Println("recv msg failed", err)
			os.Exit(1)
		}
		fmt.Println("recv msg:", string(buf[:n]))
	}
}

func main() {
	conn, err := dailServer(":8080")
	if err != nil {
		fmt.Println("connect server failed ", err)
		return
	}
	defer conn.Close()
	go recvMsg(conn)

	reader := bufio.NewReader(os.Stdin)

	for {
		str, err := reader.ReadString('\n')
		if err != nil {
			fmt.Println(err)
			continue
		}
		conn.Write([]byte(str))
	}
}
