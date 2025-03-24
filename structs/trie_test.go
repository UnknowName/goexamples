package structs

import (
    "log"
    "testing"
)

func TestTrie_Insert(t *testing.T) {
    ss := []string{"abc", "abcde", "ab", "def", "cbf"}
    tree := NewTrie()
    for _, s := range ss {
        tree.Insert(s)
    }
    // log.Println(tree.Pass)
    log.Println(tree.Search("xabcd"))
    log.Println(tree.PreCnt("abc"))
}
