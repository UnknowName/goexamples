package queue

import "fmt"

func LetterCombinations(digits string) []string {
    re := make([]string, 0)
    alphs := []string{"abc", "def", "ghi", "jkl", "mno", "pqrs", "tuv", "wxyz"}
    for i := range digits {
        index := digits[i] - '0' - 2
        for _, x := range alphs[index] {
            fmt.Println(x)
        }
    }
    return re
}

