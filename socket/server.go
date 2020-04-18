package socket

import (
	"fmt"
	"io"
	"net"
)

func TCPServer(addr string) {
	listener, err := net.Listen("tcp", addr)
	if err != nil {
		fmt.Println("Listen on 127.0.0.1:8899 error ", err)
		return
	}
	defer listener.Close()
	for {
		conn, err := listener.Accept()
		if err != nil {
			fmt.Println("Accept from client error ", err)
			return
		}
		// go handler(conn)
		echo := NewEchoHandler()
		myConn := NewConnection(conn, 1, echo)
		go myConn.Start()
	}
}

func handler(conn net.Conn) {
	defer conn.Close()
	for {
		buf := make([]byte, 1024)
		_, err := conn.Read(buf)
		if err != nil && err == io.EOF {
			fmt.Println("Client send over")
			break
		}
		fmt.Println("Received msg is ", string(buf))
	}
}

/* 定义一个Handler的Interface*/
type Handler interface {
	// 主函数前
	Prepare(conn net.Conn)
	// 主处理函数
	Execute(conn net.Conn)
	// 主函数后
	After(conn net.Conn)
}

// Echo回显的一个实现
type EchoHandler struct {}

func NewEchoHandler() Handler {
	return &EchoHandler{}
}

func (e *EchoHandler) Prepare(conn net.Conn) {
	fmt.Println("执行处理前函数...")
	if _, err := conn.Write([]byte("MESSAGE FROM SERVER")); err != nil {
		fmt.Println("向客户端写入失败", err)
	}
	fmt.Println("写入成功")
}

func (e *EchoHandler) Execute(conn net.Conn) {
	fmt.Println("执行处理主函数...")
	clientData := make([]byte, 20)
	// 一直从客户端中读出数据，读出什么，返回什么
	for {
		buf := make([]byte, 20)
		n, err := conn.Read(buf)
		if n == 0 || err == io.EOF {
			fmt.Println("Received from client over")
			break
		}
		fmt.Println("Received message from client ", string(buf))
		clientData = append(clientData, buf...)

		// 用户发送完毕，将用户发送的数据写回给客户端
		_, err = conn.Write(clientData)
		if err != nil {
			fmt.Println("Send msg to client failed ", err)
		}
	}
	_, err := conn.Write([]byte("test new msg"))
	if err != nil {
		fmt.Println("Send error ", err)
	}
}

func (e *EchoHandler) After(conn net.Conn) {
	fmt.Println("执行处理后函数...")
}

/* 将客户端的连接请教封装为自定义的Connection */
type Connection struct {
	// 底层Socket
	conn net.Conn
	// 自定义的连接ID
	id uint32
	// 连接绑定的处理方法
	handler Handler
}

func (c *Connection) Start() {
	c.handler.Prepare(c.conn)
	c.handler.Execute(c.conn)
	c.handler.After(c.conn)
}

func NewConnection(conn net.Conn, id uint32, handler Handler)  *Connection{
	return &Connection{
		conn:    conn,
		id:      id,
		handler: handler,
	}
}

/*
// 1. 用户请求的数据封装成自定义的Message格式，解决TCP粘包问题
type Message struct {
	msgId      uint32
	dataLength uint32
	data       []byte
}

func (m *Message) GetHeaderLength() uint32 {
	return 4
}

func (m *Message) MessageID() uint32 {
	return m.msgId
}

func (m *Message) DataLength() uint32 {
	return m.dataLength
}

func (m *Message) Data() []byte {
	return m.data
}

// 将原始数据打包成自定义Message
type (m *Message) Pack() []byte {
	return nil
}

// 解包
type (m *Message) UnPack() Message {
	return nil
}

// 2. 用户连接请求封装成Request, Request中包括原始socket与用户请求的数据
type Request struct {
	conn net.Conn
	Message
}

type (r *Request) GetConnect() net.Conn {
	return r.conn
}
 */