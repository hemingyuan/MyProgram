package main

import (
	"fmt"
	"net"
	"strings"

	"github.com/astaxie/beego/logs"
)

func listenServer(addr string) (l net.Listener, err error) {
	l, err = net.Listen("tcp", addr)
	return
}

func runServer(l net.Listener) {
	for {
		conn, err := l.Accept()
		if err != nil {
			logs.Warning("accept conn from [%v] failed, %v", l.Addr, err)
			conn.Close()
			continue
		}

		clientMgr.lock.Lock()
		clientMgr.clientMap[conn] = true
		clientMgr.lock.Unlock()
		fmt.Println("accept new conn", conn.RemoteAddr())
		fmt.Printf("%#v\n", clientMgr.clientMap)
		go process(conn)
	}
}

func process(conn net.Conn) {
	defer func() {
		fmt.Println("in close conn defer...")
		clientMgr.lock.Lock()
		delete(clientMgr.clientMap, conn)
		clientMgr.lock.Unlock()
	}()
	defer conn.Close()

	for {
		buf := make([]byte, 1024)
		n, err := conn.Read(buf)
		if err != nil {
			fmt.Println(err)
			return
		}
		res := string(buf[:n])
		res = strings.TrimSpace(res)
		if len(res) == 0 {
			continue
		}
		clientMgr.addMsg(conn, []byte(res))
	}
}
