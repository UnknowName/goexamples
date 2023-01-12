package main

import (
    "fmt"
    "time"
)

func main() {
    ticker := time.NewTicker(time.Second)
    for v := range ticker.C {
        fmt.Println("v is ", v)
    }
}
