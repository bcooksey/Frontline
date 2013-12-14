package engine

import "testing"

func TestDice(t *testing.T) {
    dice := Dice{sides: 6}

    roll := dice.Roll()
    if roll < 1 || roll > 6 {
        t.Errorf("Dice rolled an invalid number (%d)", roll)
    }

    // battle.
}
