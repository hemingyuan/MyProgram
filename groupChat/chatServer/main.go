package main

import (
	"fmt"

	"github.com/astaxie/beego/logs"
)

var clientMgr *ClientMgr

func init() {
	clientMgr = NewClientMgr()
}

func main() {
	logs.SetLevel(logs.LevelDebug)

	l, err := listenServer(":8080")
	if err != nil {
		logs.Error("listen server failed", err)
		return
	}
	logs.Debug("server start...")
	runServer(l)
	fmt.Println("server exited...")
}
