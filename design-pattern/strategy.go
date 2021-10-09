package design_pattern

import "log"

/*
策略模式
适用场景:
1. 当你想使用对象中各种不同的算法变体，并希望能在运行时切换算法
2. 当你有许多仅在执行某些行为时略有不同的相似类时
3. 如果算法在上下文逻辑中不是特别重要
4. 当类中使用了复杂条件运算符以同一算法的不同变体中切换
*/

// 模拟Redis的缓存清除算法，LRU/FIFO/LFU

type cleanAlgo interface {
	clean(c *cache)
}

type lruAlgo struct {
	name string
}

func (l *lruAlgo) clean(c *cache) {
	log.Println("use", l.name, "algo clean key")
}

type lfuAlgo struct {
    name string
}

func (fa *lfuAlgo) clean(c *cache) {
    log.Println("use", fa.name, "algo clean key")
}

func NewCache(max int, algo cleanAlgo) *cache {
	return &cache{
		data:    make(map[string]interface{}),
		max:     max,
		current: 0,
		algo:    algo,
	}
}

type cache struct {
	data    map[string]interface{}
	max     int       // 最大数
	current int       // 当前已使用
	algo    cleanAlgo // 当前cache使用的清除算法
}

func (c *cache) ChangeAlgo(algo cleanAlgo) {
	c.algo = algo
}

func (c *cache) AddKey(key string, value interface{}) {
    if c.current >= c.max {
        c.algo.clean(c)
        c.current--
    }
    c.data[key] = value
    c.current++
}
