package design_pattern

import (
    "log"
    "testing"
)

func TestNewSaleMachine(t *testing.T) {
    var err error
    sale := NewSaleMachine(0, 10.2)
    err = sale.AddItem(10)
    if err != nil {
        log.Fatalln(err)
    }
    err = sale.RequestItem()
    if err != nil {
        log.Fatalln(err)
    }
    err = sale.InsertMoney(15)
    if err != nil {
        log.Fatalln(err)
    }
    err = sale.DispenseItem()
    if err != nil {
        log.Fatalln(err)
    }
}
