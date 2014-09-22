package engine

import "testing"
import "fmt"

var _ = fmt.Println // TODO: Delete

/****** Research Phase ******/

func TestBuyAttempts(t *testing.T) {
    phase := ResearchPhase{Phase{step: 1}}

    if phase.BuyAttempts(1, 5) != true {
        t.Error("Research Phase - Couldn't buy attempts during first step")
    }

    if phase.CurrentStep() != 2 {
        t.Error("Research Phase - Did not advance from step 1 to 2")
    }
}

func TestBuyAttemptsFailsWhenSuppliesTooLow(t *testing.T) {
    phase := ResearchPhase{Phase{step: 1}}
    // turn := Turn{}

    if phase.BuyAttempts(2, 5) != false {
        t.Error("Research Phase - Bought attempts player could not afford")
    }

    if phase.CurrentStep() != 1 {
        t.Error("Research Phase - Wrongly advanced to next step when purchase failed")
    }
}
