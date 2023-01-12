package main

import (
    "fmt"
    "testing"
    "time"
)

func TestDoWork(t *testing.T) {
    done := make(chan interface{})
    time.AfterFunc(time.Second * 5, func() {
        close(done)
    })
    heart, results := DoWork(done,time.Second)
    for {
        select {
        case _, ok := <- heart:
            if !ok {
                return
            }
            fmt.Println("pulse...")
        case r, ok := <- results:
            if !ok {
                return
            }
            fmt.Printf("results %v\n", r.Second())
        case <-time.After(time.Second * 2):
            return
        }
    }
}
