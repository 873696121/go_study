package main

import (
	"net"
	"strings"
)

type User struct {
	Name string
	Addr string
	C    chan string
	conn net.Conn

	Server *Server
}

// 创建一个用户的API
func NewUser(conn net.Conn, server *Server) *User {
	userAddr := conn.RemoteAddr().String()

	user := &User{
		Name:   userAddr,
		Addr:   userAddr,
		C:      make(chan string),
		conn:   conn,
		Server: server,
	}

	go user.ListenMessage()

	return user
}

// 用户上线业务
func (this *User) Online() {
	//用户上线，将用户加入到map中
	this.Server.mapLock.Lock()
	this.Server.OnlineMap[this.Name] = this
	this.Server.mapLock.Unlock()

	//广播消息
	this.Server.BroadCast(this, "已上线")
}

// 用户下线业务
func (this *User) Offline() {
	//用户上线，将用户加入到map中
	this.Server.mapLock.Lock()
	delete(this.Server.OnlineMap, this.Name)
	this.Server.mapLock.Unlock()

	//广播消息
	this.Server.BroadCast(this, "下线")
}

// 给当前user对应的客户端发消息
func (this *User) SendMsg(msg string) {
	this.conn.Write([]byte(msg))
}

// 用户处理消息的业务
func (this *User) DoMessage(msg string) {
	if msg == "who" {
		// 查询当前在线用户有哪些
		this.Server.mapLock.Lock()
		for _, user := range this.Server.OnlineMap {
			onlineMessage := "[" + user.Addr + "]" + user.Name + ":" + "在线...\n"
			this.SendMsg(onlineMessage)
		}
		this.Server.mapLock.Unlock()
	} else if len(msg) > 7 && msg[:7] == "rename|" {
		newName := strings.Split(msg, "|")[1]

		_, ok := this.Server.OnlineMap[newName]
		if ok {
			this.SendMsg("当前用户名被使用\n")
		} else {
			this.Server.mapLock.Lock()
			delete(this.Server.OnlineMap, this.Name)
			this.Server.OnlineMap[newName] = this
			this.Server.mapLock.Unlock()

			this.Name = newName
			this.SendMsg("您已经更新用户名:" + this.Name + "\n")
		}
	} else if len(msg) > 4 && msg[:3] == "to|" {
		// to|zhangsan|content
		// 1. 获取对方用户名
		remoteName := strings.Split(msg, "|")[1]

		if remoteName == "" {
			this.SendMsg("消息格式不正确，请使用 to|zhangsan|content \n")
			return
		}

		// 2. 根据用户名得到用户对象
		remoteUser, ok := this.Server.OnlineMap[remoteName]
		if !ok {
			this.SendMsg("该用户名不存在\n")
			return
		}

		// 3. 获取消息内容，通过对方user对象传过去
		content := strings.Split(msg, "|")[2]
		if content == "" {
			this.SendMsg("无消息内容，请重发")
			return
		}

		remoteUser.SendMsg(this.Name + "对您说: " + content + "\n")

	} else {
		this.Server.BroadCast(this, msg)
	}
}

// 监听当前user channel的方法, 一旦有消息， 就直接发送给客户端
func (this *User) ListenMessage() {
	for {
		msg := <-this.C
		this.conn.Write([]byte(msg + "\n"))
	}
}
