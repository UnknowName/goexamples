package socket

import (
	"fmt"
	"io"
	"net"
	"time"
)

func Telnet() {
	conn, err := net.Dial("tcp", "127.0.0.1:8899")
	if err != nil {
		fmt.Println("Connect error ", err)
		return
	}
	defer conn.Close()
	for i := 0; i < 2; i++ {
		msg := fmt.Sprintf("hello,world %d", i)
		_, err = conn.Write([]byte(msg))
	}
	for {
		buf := make([]byte, 20)
		// 读动作会一直阻塞，直到从conn读到数据
		n, err := conn.Read(buf)
		if n == 0 && err == io.EOF {
			fmt.Println("Receive from server over")
			break
		}
		if err != nil {
			/* 发生错误，一般意味着要进行重连 */
			fmt.Println("No data received from server , error is ", err)
			for {
				conn, err = reConnect("tcp", "127.0.0.1:8899")
				if err == nil {
					break
				}
				time.Sleep(time.Second)
			}
		}
		fmt.Println("Received message from server ", string(buf))
	}
}

func reConnect(network, address string) (net.Conn, error) {
	conn, err := net.Dial(network, address)
	return conn, err
}
