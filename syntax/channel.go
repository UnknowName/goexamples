package syntax

import (
	"fmt"
	"net/http"
	"time"
)

type Site struct {
	Url   string
	Host  string
	Status int
}

func Check(s *Site) {
	client := &http.Client{Timeout: 5 * time.Second}
	req, _ := http.NewRequest("GET", s.Url, nil)
	req.Host = s.Host
	resp, err := client.Do(req)
	if err != nil {
		fmt.Println(err)
		return
	}
	defer resp.Body.Close()
	s.Status = resp.StatusCode
	ch := make(chan *Site, 1)
	// 在往Channel放入数据前，要先准备好从通道取数据的协程，不然主协程会一直阻塞
	go Count(ch)
	ch <- s
	fmt.Println("Check func Over")
}

func Count(ch chan *Site) {
	s := <- ch
	fmt.Println("Get site struct from ch", s.Status)
}