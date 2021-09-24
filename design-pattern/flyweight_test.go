package design_pattern

import (
    "log"
    "testing"
)

func TestNewPlayer(t *testing.T) {
    players := make([]*player, 0)
    players = append(players, NewPlayer("red"))
    players = append(players, NewPlayer("red"))
    players = append(players, NewPlayer("red"))
    players = append(players, NewPlayer("red"))
    players = append(players, NewPlayer("red"))
    // 五个play.address值相同，因为共享了底层的结构体
    for _, play := range players {
        log.Println(play.address)
    }
}