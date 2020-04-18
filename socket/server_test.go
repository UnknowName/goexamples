package socket

import "testing"

func TestTCPServer(t *testing.T) {
	addr := "127.0.0.1:8899"
	go TCPServer(addr)
	add2 := "127.0.0.1:9988"
	go TCPServer(add2)
	select {}
}

