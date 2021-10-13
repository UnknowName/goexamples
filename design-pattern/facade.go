package design_pattern

import (
    "log"
)

/*
门面模式/外观模式: 为程序库、 框架或其他复杂类提供一个简单的接口。外观模式一般同单例一同使用
适用场景
1. 如果需要一个指向复杂子系统的直接接口，且该接口功能有限
2. 如果需要将子系统组织成多层结构
 */

// 模拟通过电话使用信用卡订购商品，实际系统会执行一系列动作
/*
1. 检查帐号
2. 检查安全码
3. 信用卡余额
4. 帐单录入
5. 发送消息通知
 */

// 复杂子系统
type account struct {
    name string
}

// 检查帐户方法
func (a *account) checkAccount(name string) bool {
    return a.name == name
}

type securityCode struct {
    code int
}

// 检查安全码
func (s *securityCode) checkCode(code int) bool {
    return s.code == code
}

type wallet struct {
    balance float64
}

// 检查余额
func (w *wallet) checkBalance(price float64) bool {
    return w.balance >= price
}

func NewWalletFacade(accountName string, code int) *walletFacade {
    log.Println("Starting create account")
    obj := &walletFacade{
        account: &account{name: accountName},
        securityCode: &securityCode{code: code},
        wallet: &wallet{balance: 0},
    }
    log.Println("Account created")
    return obj
}

// 外观将复杂子系统封装进来
type walletFacade struct {
    account *account
    securityCode *securityCode
    wallet *wallet
}

// 外观类实现的方法，客户端只与它交互，这样不同的子系统就各司其职

func (wf *walletFacade) Order(accountName string, price float64) {
    log.Println("start order ")
    if wf.account.checkAccount(accountName) && wf.wallet.checkBalance(price) {
        log.Println("check ok")
    } else {
        log.Println("check failed")
    }
    log.Println("end order")
}