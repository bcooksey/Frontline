package engine

// TODO: Need to decide on interface for a phase. maybe:
//   UpdateTurnFromPhase (for applying new research, purchased units, etc.)
//   NextPhase() - Returns next phase so easy to continue driving through states

import "fmt"

var _ = fmt.Println // TOOD: DELETE

type ResearchPhase struct {
    state ResearchState
}

func (phase *ResearchPhase) Init(startAt int) {
    if startAt == 0 {
        phase.state = &ResearchNoAttemptsState{Phase: phase}
    } else if startAt == 1 {
        phase.state = &ResearchAttemptsBoughtState{Phase: phase}
    }
}

func (phase *ResearchPhase) BuyAttempts(numberOfAttempts int, availableSupplies int) bool {
    return phase.state.BuyAttempts(numberOfAttempts, availableSupplies)
}

func (phase *ResearchPhase) AttemptResearch(numberOfAttempts int, targetCategory int, dice Roller) bool {
    return phase.state.AttemptResearch(numberOfAttempts, targetCategory, dice)
}

func (phase ResearchPhase) GetState() ResearchState {
    return phase.state
}

func (phase *ResearchPhase) SetState(state ResearchState) {
    phase.state = state
}

type ResearchState interface {
    BuyAttempts(int, int) bool
    AttemptResearch(int, int, Roller) bool
}

type ResearchNoAttemptsState struct {
    Phase *ResearchPhase
}

func (state *ResearchNoAttemptsState) BuyAttempts(numberOfAttempts int, availableSupplies int) bool {
    if numberOfAttempts*5 <= availableSupplies {
        state.Phase.SetState(&ResearchAttemptsBoughtState{Phase: state.Phase})
        return true
    }
    return false
}

func (state *ResearchNoAttemptsState) AttemptResearch(numberOfAttempts int, targetCategory int, dice Roller) bool {
    return false
}

type ResearchAttemptsBoughtState struct {
    Phase *ResearchPhase
}

func (state *ResearchAttemptsBoughtState) BuyAttempts(numberOfAttempts int, availableSupplies int) bool {
    return false
}

func (state *ResearchAttemptsBoughtState) AttemptResearch(numberOfAttempts int, targetCategory int, dice Roller) bool {
    if numberOfAttempts > 0 {
        for i := 0; i < numberOfAttempts; i++ {
            roll := dice.Roll()
            if roll == targetCategory {
                state.Phase.SetState(&ResearchNoAttemptsState{Phase: state.Phase})
                return true
            }
        }
    }
    state.Phase.SetState(&ResearchNoAttemptsState{Phase: state.Phase})
    return false
}
