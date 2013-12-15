package engine

import "testing"

type Mock struct{}

func (m *Mock) Attack(int) bool  { return true }
func (m *Mock) Defend(int) bool  { return true }
func (m *Mock) Category() string { return "land" }

func TestCreateBattle(t *testing.T) {
    battle := CreateBattle(nil, nil, nil)

    if battle.Phase() != "attack" {
        t.Error("New Battle starts in wrong phase")
    }

    if battle.dice == nil {
        t.Error("New Battle does not automatically create a die")
    }
}

func TestBattleAttacking(t *testing.T) {
    attackers := make([]Attacker, 2)
    attackers[0] = &Mock{}
    attackers[1] = &Mock{}

    battle := CreateBattle(attackers, nil, nil)

    hits := battle.RollForAttackers()
    count, ok := hits["land"]
    if !ok {
        t.Error("Rolling for attackers does not report hit count for land units")
    }
    if count != 2 {
        t.Error("Hit count returned an invalid value")
    }
}

func TestBattleDefending(t *testing.T) {
    defenders := make([]Defender, 2)
    defenders[0] = &Mock{}
    defenders[1] = &Mock{}

    battle := CreateBattle(nil, defenders, nil)

    hits := battle.RollForDefenders()
    count, ok := hits["land"]
    if !ok {
        t.Error("Rolling for defenders does not report hit count for land units")
    }
    if count != 2 {
        t.Error("Hit count returned an invalid value")
    }
}
