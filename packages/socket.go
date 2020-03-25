package packages

import (
	"bufio"
	"fmt"
	"net"
	"time"
)

func EchoServer() {
	listener, err := net.Listen("tcp", "localhost:8080")
	failed(err, "Listen on failed")
	defer listener.Close()
	for {
		conn, err := listener.Accept()
		failed(err, "accept failed")
		go handler(conn)
	}
}

func failed(err error, msg string) {
	if err !=nil {
		fmt.Println(msg, err)
		return
	}
}

func handler(c net.Conn) {
	_, _ = c.Write([]byte("Welcome to GoEcho\n"))
	defer c.Close()
	for {
		reader := bufio.NewReader(c)
		s , err := reader.ReadString('\n')
		if err != nil {
			fmt.Println("Read from socket err ", err)
			return
		}
		fmt.Println(s)
		_, err = c.Write([]byte(s))
		if err != nil {
			return
		}
		time.Sleep(1 * time.Second)
	}
}
