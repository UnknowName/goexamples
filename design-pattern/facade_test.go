package design_pattern

import (
    "testing"
)

func TestNewWalletFacade(t *testing.T) {
    // 直接使用外观模式，将创建帐号/检查帐户等一系统工作交给外观类实现
    // 实际外观又是调用实在类处理这些事情
    facade := NewWalletFacade("alice", 123)
    facade.Order("blob", 12.1)
}
