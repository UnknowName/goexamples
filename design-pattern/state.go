package design_pattern

import (
    "errors"
    "log"
)

/*
状态模式
适用场景:
1. 如果对象需要根据自身当前状态进行不同行为， 同时状态的数量非常多且与状态相关的代码会频繁变更的话
2. 如果某个类需要根据成员变量的当前值改变自身行为， 从而需要使用大量的条件语句时， 可使用该模式
3. 当相似状态和基于条件的状态机转换中存在许多重复代码时， 可使用状态模式
*/

// 定义状态接口,每个状态都要实现该接口

type state interface {
	addItem(int) error
	requestItem() error
	insertMoney(int) error
	dispenseItem() error
}

// 具体状态

// 有商品

type hasItemState struct {
    saleMachine *saleMachine
}

func (hs *hasItemState) addItem(n int) error {
    hs.saleMachine.itemCount += n
    return nil
}

// 有商品的情况下，请求商品后更改状态

func (hs *hasItemState) requestItem() error {
    if hs.saleMachine.itemCount > 0 {
        // 商品够，进入商品请求阶段
        hs.saleMachine.currentState = hs.saleMachine.itemRequested
    } else {
        // 商品不足，进入缺少状态
        hs.saleMachine.ChangeState(hs.saleMachine.noItem)
    }
    return nil
}

// 当前状态下不允许的操作
func (hs *hasItemState) insertMoney(n int) error {
    return errors.New("please select item first")
}

func (hs *hasItemState) dispenseItem() error {
    return errors.New("please select item first")
}

// 无商品

type noItemState struct {
    saleMachine *saleMachine
}

func (ns *noItemState) addItem(n int) error {
    ns.saleMachine.itemCount += n
    ns.saleMachine.currentState = ns.saleMachine.hasItem
    return nil
}

func (ns *noItemState) dispenseItem() error {
    return errors.New("no item to sale")
}

func (ns *noItemState) insertMoney(n int) error {
    return errors.New("no item to sale")
}

func (ns *noItemState) requestItem() error {
    return errors.New("no item to sale")
}


// 商品已请求

type itemRequestedState struct {
    saleMachine *saleMachine
}

func (is *itemRequestedState) addItem(n int) error {
    return errors.New("please insert money")
}

func (is *itemRequestedState) requestItem() error {
    return errors.New("has selected item")
}

func (is *itemRequestedState) insertMoney(n int) error {
    if float64(n) >= is.saleMachine.itemPrice {
        is.saleMachine.currentState = is.saleMachine.hasMoney
        is.saleMachine.itemCount -= 1
    } else {
        return errors.New("less money")
    }
    return nil
}

func (is *itemRequestedState) dispenseItem() error {
    return errors.New("please insert money first")
}

// 收到纸币

type hasMoneyState struct {
    saleMachine *saleMachine
}

func (hm *hasMoneyState) dispenseItem() error {
    log.Println("Dispensing Item")
    hm.saleMachine.itemCount -= 1
    if hm.saleMachine.itemCount > 0 {
        hm.saleMachine.currentState = hm.saleMachine.hasItem
    } else {
        hm.saleMachine.currentState = hm.saleMachine.noItem
    }
    log.Println("sale success")
    return nil
}

func (hm *hasMoneyState) insertMoney(n int) error {
    return errors.New("has inserted money")
}

func (hm *hasMoneyState) requestItem() error {
    return errors.New("has selected item")
}

func (hm *hasMoneyState) addItem(n int) error {
    return errors.New("item dispense in progress")
}

// 自动售货机

func NewSaleMachine(total int, price float64) *saleMachine {
    s := &saleMachine{}
    s.hasItem = &hasItemState{saleMachine: s}
    s.noItem = &noItemState{saleMachine: s}
    s.itemRequested = &itemRequestedState{saleMachine: s}
    s.hasMoney = &hasMoneyState{saleMachine: s}
    s.itemPrice = price
    s.itemCount = total
    if total > 0 {
        s.currentState = s.hasItem
    } else {
        s.currentState = s.noItem
    }
    return s
}

// 它也要实现状态的接口

type saleMachine struct {
	hasItem       state
	noItem        state
	itemRequested state
	hasMoney      state
	// currentState保存着当前状态，当调用不同方法时，方法判断是否应该转变状态
	currentState  state
	itemCount     int
	itemPrice     float64
}

func (sm *saleMachine) ChangeState(newState state) {
    sm.currentState = newState
}

func (sm *saleMachine) AddItem(n int) error {
    return sm.currentState.addItem(n)
}

func (sm *saleMachine) RequestItem() error {
    return sm.currentState.requestItem()
}

func (sm *saleMachine) InsertMoney(n int) error {
    return sm.currentState.insertMoney(n)
}

func (sm *saleMachine) DispenseItem() error {
    return sm.currentState.dispenseItem()
}