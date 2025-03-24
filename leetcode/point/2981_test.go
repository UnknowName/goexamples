package point

import (
    "log"
    "testing"
)

func TestMaximumLength(t *testing.T) {
    s := "aaaa"
    n := MaximumLength(s)
    log.Println(n)
}
