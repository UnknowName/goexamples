package datastruct

import (
	"fmt"
	"testing"
)

func TestLineLink_Insert(t *testing.T) {
	link := NewLineLink()
	elem := Element{Name: "a", ID: 123}
	link.Insert(elem, 1)
	elem2 := Element{Name: "b", ID: 124}
	link.Insert(elem2, 2)
	elem3 := Element{Name: "b", ID: 125}
	link.Insert(elem3, 1)
	fmt.Println("Before delete ", link)
	link.Delete(1)
	fmt.Println("After delete ", link)
	fmt.Println(len("abc"))
}

func TestLineLink_Delete(t *testing.T) {
	equeal("abcd", "c")
}

func equeal(s1, s2 string) {
	i, j := 0, 0
	len1, len2 := len(s1), len(s2)
	fmt.Println(len1, len2)
	for i < len1 && j < len2 {
		if s1[i] == s2[j] {
			i += 1
			j += 1
		} else {
			i = i - j + 1
			j = 0
		}
		fmt.Println("i is ", i, "j is ", j)
		//time.Sleep(time.Second)
	}
	if j == len2 {
		fmt.Println("True")
		//os.Exit(0)
	} else {
		fmt.Println("false")
	}
}
