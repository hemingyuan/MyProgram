package main

import (
	"errors"
	"fmt"
	"net"
	"sync"
	"time"

	"github.com/astaxie/beego/logs"
)

type sendClient struct {
	msg  []byte
	conn net.Conn
}

// ClientMgr manage client connect
type ClientMgr struct {
	clientMap map[net.Conn]bool
	MsgChan   chan *sendClient
	lock      sync.RWMutex
}

// NewClientMgr init function
func NewClientMgr() *ClientMgr {
	mgr := &ClientMgr{
		clientMap: make(map[net.Conn]bool),
		MsgChan:   make(chan *sendClient, 1024),
	}
	go mgr.broadcast()
	return mgr
}

func (c *ClientMgr) broadcast() {
	for msg := range c.MsgChan {
		c.forward(msg)
	}
}

func (c *ClientMgr) forward(msg *sendClient) {
	c.lock.RLock()
	for conn := range c.clientMap {
		if conn == msg.conn {
			continue
		}
		c.sendMsgToCli(conn, msg.msg)
	}
	c.lock.RUnlock()
}

func (c *ClientMgr) sendMsgToCli(conn net.Conn, msg []byte) {
	msgLen := len(msg)
	pos := 0

	for pos < msgLen {
		n, err := conn.Write(msg)
		if err != nil {
			logs.Warning("send msg to [%v] failed,err:%v", conn.RemoteAddr(), err)
			conn.Close()
			delete(c.clientMap, conn)
			return
		}
		pos += n
		fmt.Println("msg:", string(msg), msgLen, pos, n)
		msg = msg[n:]
	}
}

func (c *ClientMgr) addMsg(conn net.Conn, msg []byte) (err error) {
	sendMsg := &sendClient{
		msg:  msg,
		conn: conn,
	}
	t := time.After(time.Second)
	select {
	case <-t:
		logs.Warning("send msg to chan timeout")
		err = errors.New("send msg to chan timeout")
	case c.MsgChan <- sendMsg:
		logs.Debug("add msg to chan success", string(msg))
	}
	return
}
