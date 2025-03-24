package datastruct

import (
    "fmt"
    "testing"
)

func TestUnion_Merge(t *testing.T) {
    union := NewUnion(10)
    fmt.Println(union.total == 10)
    union.Merge(0, 9)
    fmt.Println(union.Find(0) == union.Find(9), union.total == 9)
}
