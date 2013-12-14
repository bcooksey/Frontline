package engine

import "testing"

type Mock struct{}

func (m *Mock) Attack(int) bool  { return true }
func (m *Mock) Category() string { return "land" }

func TestBattle(t *testing.T) {
    attackers := make([]Attacker, 2)
    attackers[0] = &Mock{}
    attackers[1] = &Mock{}

    defenders := make([]Defender, 1)

    battle := CreateBattle(attackers, defenders, nil)

    if battle.Phase() != "attack" {
        t.Error("New Battle starts in wrong phase")
    }

    hits := battle.RollForAttackers()
    count, ok := hits["land"]
    if !ok {
        t.Error("Rolling for attackers does not report hit count for land units")
    }
    if count != 2 {
        t.Error("Hit count returned an invalid value")
    }
}
