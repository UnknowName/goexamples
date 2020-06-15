package algo

import (
	"fmt"
	"testing"
)

func TestBubbleSort(t *testing.T) {
	lst := []int{90, 1, 10, -1, 1, 5, 30}
	BubbleSort(lst)
	fmt.Println(lst)
}
