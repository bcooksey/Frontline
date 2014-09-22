package engine

import "testing"

type MockDice struct{}

func (m *MockDice) Roll() int { return 1 }

/****** Research Phase ******/

func TestBuyAttempts(t *testing.T) {
    phase := ResearchPhase{Phase{step: 1}}

    if phase.BuyAttempts(1, 5) != true {
        t.Error("Research Phase - Couldn't buy attempts during first step")
    }

    if phase.CurrentStep() != 2 {
        t.Error("Research Phase - Did not advance to next step")
    }

    if phase.BuyAttempts(1, 5) != false {
        t.Error("Research Phase - Bought attempts during an incorrect step")
    }
}

func TestBuyAttemptsFailsWhenSuppliesTooLow(t *testing.T) {
    phase := ResearchPhase{Phase{step: 1}}

    if phase.BuyAttempts(2, 5) != false {
        t.Error("Research Phase - Bought attempts player could not afford")
    }

    if phase.CurrentStep() != 1 {
        t.Error("Research Phase - Wrongly advanced to next step when purchase failed")
    }
}

func TestAttemptResearch(t *testing.T) {
    phase := ResearchPhase{Phase{step: 2}}
    dice := &MockDice{}

    if phase.AttemptResearch(1, dice) != true {
        t.Error("Research Phase - Couldn't attempt research during first step")
    }

    if phase.CurrentStep() != 3 {
        t.Error("Research Phase - Did not advance to next step")
    }
}
