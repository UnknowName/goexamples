package algo

import (
    "fmt"
    "log"
)

func NewRingBuffer(size int) *RingBuffer {
    return &RingBuffer{
        data: make([]byte, size),
        rIndex: 0,
        wIndex: 0,
        size: 0,
    }
}

type RingBuffer struct {
    data []byte
    rIndex int
    wIndex int
    size   int
}

func (rb *RingBuffer) String() string {
    return fmt.Sprintf("RingBuf(%v,r=%d,w=%d,size=%d)",
        rb.data,rb.rIndex,rb.wIndex,rb.size)
}

func (rb *RingBuffer) IsEmpty() bool {
    return rb.size == 0
}

func (rb *RingBuffer) IsFull() bool {
    return rb.size == len(rb.data)
}

func (rb *RingBuffer) Add(n byte) {
    if rb.IsFull() {
        log.Println("full ring buffer")
        return
    }
    rb.data[rb.wIndex] = n
    rb.wIndex = (rb.wIndex + 1) % len(rb.data)
    rb.size++
}

func (rb *RingBuffer) Get(length int) []byte {
    if rb.IsEmpty() {
        log.Println("empty")
        return nil
    }
    var buf []byte
    // 取的长度大于队列的长度，则以实际为准
    if length > rb.size {
        length = rb.size
    }
    end := rb.rIndex + length
    if end > len(rb.data) {
        // 先取尾部
        buf = rb.data[rb.rIndex:]
        // 再从头部继续取
        remain := length - len(buf)
        index := remain % len(rb.data)
        buf = append(buf, rb.data[:index]...)
        rb.size -= length
        rb.rIndex = index
    } else {
        // 直接取，退化成队列
        buf = rb.data[rb.rIndex:end]
        rb.size -= length
        rb.rIndex = (rb.rIndex + length) % len(rb.data)
    }
    return buf
}