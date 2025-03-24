package tmp

import (
    "encoding/json"
    "fmt"
    "log"
    "strings"
    "testing"
)

func TestVector2Bytes(t *testing.T) {
    vector := []float64{1.0,2.0,3.0,4.0,5.0}
    s := fmt.Sprintf("%v", vector)
    v := strings.Replace(s, " ", ",", -1)
    log.Println(v)
    var nums []int
    if err := json.Unmarshal([]byte(v), &nums); err != nil {
        log.Println(err)
    }
    log.Printf("nums is %v", nums)
}


func TestGetBin(t *testing.T) {
}

func TestCalDict(t *testing.T) {
    nums := []int{2,4,6,8}
    v := minSizeSubarray(nums, 3)
    fmt.Println(v)
}