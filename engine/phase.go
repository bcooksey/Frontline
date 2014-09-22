package engine

type Phase struct {
    step int
}

func (phase Phase) CurrentStep() int { return phase.step }

type ResearchPhase struct {
    Phase
}

func (phase *ResearchPhase) BuyAttempts(numberOfAttempts int, availableSupplies int) bool {
    if numberOfAttempts * 5 <= availableSupplies {
        phase.step = 2
        return true
    }
    return false
}
