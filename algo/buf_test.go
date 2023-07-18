package algo

import (
    "log"
    "testing"
)

func TestRingBuffer_Add(t *testing.T) {
    buf := NewRingBuffer(5)
    var i byte
    i = 1
    for i < 6 {
        buf.Add(i)
        i++
    }
    res := buf.Get(2)
    log.Println(res)
    buf.Add(6)
    buf.Add(7)
    res = buf.Get(5)
    log.Println(res)
}
