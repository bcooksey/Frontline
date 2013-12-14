package engine

import "testing"

func TestUnitAttack(t *testing.T) {
    u := Unit{attackRating:  2}

    if u.Attack(3) != false {
        t.Error("Unit attacked with a die roll higher than its rating")
    }

    if u.Attack(2) != true {
        t.Error("Unit did not attack with a die roll equal to its rating")
    }

    if u.Attack(1) != true {
        t.Error("Unit did not attack with a die roll lower to its rating")
    }
}
