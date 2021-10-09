package design_pattern

import (
    "testing"
)

func TestNewCache(t *testing.T) {
    algo1 := &lruAlgo{"LRU"}
    algo2 := &lfuAlgo{"LFU"}
    cache := NewCache(2, algo1)
    cache.AddKey("name1", "cheng")
    cache.AddKey("name2", "cheng")
    cache.AddKey("name3", "cheng")
    // 运行时改变清理缓存算法
    cache.ChangeAlgo(algo2)
    cache.AddKey("name4", "cheng")
}
