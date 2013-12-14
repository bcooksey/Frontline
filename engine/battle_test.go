package engine

import "testing"

func TestBattle(t *testing.T) {
    attackers := make([]Unit, 3)
    defenders := make([]Unit, 3)
    battle := CreateBattle(attackers, defenders)

    if battle.Phase() != "attack" {
        t.Error("New Battle starts in wrong phase")
    }
}
