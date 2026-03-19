package main

import (
	"fmt"
	"net"
)

type Server struct {
	// 服务器地址
	Ip string
	// 服务器端口
	Port int
}

// 创建一个sever的接口
func NewServer(ip string, port int) *Server {
	server := &Server{
		Ip:   ip,
		Port: port,
	}
	return server
}

func (this *Server) Handler(conn net.Conn) {
	// 当前连接的业务
	fmt.Println("连接建立成功")
}

// 启动服务器的接口
func (this *Server) Start() {
	//socket listen
	listener, err := net.Listen("tcp", fmt.Sprintf("%s:%d", this.Ip, this.Port))
	if err != nil {
		fmt.Println("net.Listen error:", err)
		return
	}
	//close listen socket
	defer listener.Close()
	for {
		//accept
		conn, err := listener.Accept()
		if err != nil {
			fmt.Println("listener accept error:", err)
			continue
		}
	    //do handler
		go this.Handler(conn)
	}
}