package structs

func NewTrie() *Trie {
    return &Trie{
        Pass: 0,
        End: 0,
        Next: make(map[int32]*Trie),
    }
}

type trieNode struct {
    pass int
    end  int
}

type Trie struct {
    Pass int
    End  int
    Next map[int32]*Trie
}

func (t *Trie) Insert(s string) {
    node := t
    for _, v  := range s {
        node.Pass++
        index := v - 'a'
        if node.Next[index] == nil {
            node.Next[index] = NewTrie()
        }
        node = node.Next[index]
    }
    node.Pass++
    node.End++
}

func (t *Trie) Search(key string) bool {
    node := t
    for _, v := range key {
        index := v - 'a'
        node = node.Next[index]
        if node == nil {
            return false
        }
    }
    return node.End > 0
}

func (t *Trie) PreCnt(s string) int {
    node := t
    for _, v := range s {
        index := v - 'a'
        node = node.Next[index]
        if node == nil {
            return 0
        }
    }
    return node.Pass
}

