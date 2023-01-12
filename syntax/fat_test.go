package main

import (
    "log"
    "testing"
    "time"
)

func TestRepeat(t *testing.T) {
    done := make(chan interface{})
    go func() {
        time.Sleep(time.Second)
        close(done)
    }()
    for v := range Take(done, Repeat(done, 1), 5) {
        log.Printf("%v", v)
    }
}
