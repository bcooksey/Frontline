package engine

import "testing"

type Mock struct{}

func (m *Mock) Attack(int) bool  { return true }
func (m *Mock) Defend(int) bool  { return true }
func (m *Mock) Wound() bool      { return true }
func (m *Mock) Wounded() bool    { return false }
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

func TestBattleWoundDefenders(t *testing.T) {
    defenders := make([]Defender, 2)
    defenders[0] = &Unit{category: "land"}
    defenders[1] = &Unit{category: "land"}

    battle := CreateBattle(nil, defenders, nil)

    casualties := map[string]int{}
    casualties["land"] = 2
    battle.WoundDefenders(casualties)

    if battle.Defenders()[0].Wounded() != true {
        t.Error("Some Defenders were not wounded that should have been")
    }

    defenders = make([]Defender, 2)
    defenders[0] = &Unit{category: "land"}
    defenders[1] = &Unit{category: "land"}

    battle = CreateBattle(nil, defenders, nil)
    casualties["land"] = 1
    battle.WoundDefenders(casualties)

    if battle.Defenders()[1].Wounded() != false {
        t.Errorf("Wounding defenders did not mark correct units as wounded.")
    }

    casualties["land"] = 100000
    _, err := battle.WoundDefenders(casualties)

    if err == nil {
        t.Errorf("Wounding more defenders than there are participating in the battle does not throw an error")
    }
}

func TestBattleWoundAttackers(t *testing.T) {
    attackers := make([]Attacker, 2)
    attackers[0] = &Unit{category: "land"}
    attackers[1] = &Unit{category: "land"}

    battle := CreateBattle(attackers, nil, nil)

    casualties := map[string]int{}
    casualties["land"] = 2
    battle.WoundAttackers(casualties)

    if battle.Attackers()[0].Wounded() != true {
        t.Error("Some Attackers were not wounded that should have been")
    }

    attackers = make([]Attacker, 2)
    attackers[0] = &Unit{category: "land"}
    attackers[1] = &Unit{category: "land"}

    battle = CreateBattle(attackers, nil, nil)
    casualties["land"] = 1
    battle.WoundAttackers(casualties)

    if battle.Attackers()[1].Wounded() != false {
        t.Errorf("Wounding attackers did not mark correct units as wounded.")
    }

    casualties["land"] = 100000
    _, err := battle.WoundAttackers(casualties)

    if err == nil {
        t.Errorf("Wounding more attackers than there are participating in the battle does not throw an error")
    }
}

func TestBattleRemoveCasualties(t *testing.T){
    attackers := make([]Attacker, 2)
    attackers[0] = &Unit{category: "land", wounded: true}
    attackers[1] = &Unit{category: "land"}

    defenders := make([]Defender, 1)
    defenders[0] = &Unit{category: "land", wounded: true}

    battle := CreateBattle(attackers, defenders, nil)

    battle.RemoveCasualties()

    if len(battle.Attackers()) != 1 {
        t.Error("Wrong number of ataccker casualties were removed from the battle")
    }

    if len(battle.Defenders()) != 0 {
        t.Error("Wrong number of defending casualties were removed from the battle")
    }
}
