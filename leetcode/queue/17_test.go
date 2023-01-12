package queue

import (
    "log"
    "testing"
)

func TestLetterCombinations(t *testing.T) {
    digits := "234"
    re := LetterCombinations(digits)
    log.Println("re =", re)
}
