package design_pattern

import (
    "fmt"
    "sync"
)

/*
享元模式
适用场景：
1. 程序需要生成数量巨大的相似对象
2. 对象中包含可抽取且能在多个对象间共享的重复状态
一般会配合工厂模式一起使用
*/

// 工厂方法函数，用于获取/存储共享的享元对象

var (
    // 这里需要作为全局变量，不然在getDress方法中，拿不到原来的对象
    dressFactory = newDressFactory()
)

func newDressFactory() *addressFactory {
    return &addressFactory{
        mutex: sync.Mutex{},
        dressMap: make(map[string]dress),
    }
}

type addressFactory struct {
    mutex sync.Mutex
    dressMap map[string]dress
}

func (af *addressFactory) getDress(addressType string) dress {
    af.mutex.Lock()
    defer af.mutex.Unlock()
    if _dress, exist := af.dressMap[addressType]; exist {
       return _dress
    }
    var _dress dress
    switch addressType {
    case "red":
        _dress = &redDress{color: "red"}
    case "green":
        _dress = &greenDress{color: "green"}
    default:
        panic("un support address type")
    }
    af.dressMap[addressType] = _dress
    return _dress
}

// 享元接口

type dress interface {
	getColor() string
}

// 具体的享元对象

type redDress struct {
    color string
}

func (rd *redDress) getColor() string {
    return rd.color
}

func (rd *redDress) String() string {
    return fmt.Sprintf("redDress{at: %p}", rd)
}

type greenDress struct {
    color string
}

func (gd *greenDress) getColor() string {
    return gd.color
}

// 玩家结构体，共享address对象

func NewPlayer(addressType string) *player {
	return &player{
	    address: dressFactory.getDress(addressType),
    }
}

type player struct {
	address  dress
	playType string
}
