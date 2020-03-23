package syntax

import (
	"testing"
	"time"
)

func TestCheck(t *testing.T) {
	s := &Site{
		Url:    "http://128.0.255.254/demo/daohao",
		Host:   "dev.siss.io",
		Status: 0,
	}
	Check(s)
	// 主进程要稍等一会儿
	time.Sleep(time.Second)
}