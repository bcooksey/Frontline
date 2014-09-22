package engine

type Phase struct {
    step int
}

// TODO: Need to decide on interface for a phase. maybe:
//   CurrentStep
//   UpdateTurnFromPhase (for applying new research, purchased units, etc.)
//   NextPhase() - Returns next phase so easy to continue driving through states
//   Missing a way to drive through a phase...

func (phase Phase) CurrentStep() int { return phase.step }

type ResearchPhase struct {
    Phase
}

func (phase *ResearchPhase) BuyAttempts(numberOfAttempts int, availableSupplies int) bool {
    if phase.step != 1 {
        return false
    } else if numberOfAttempts * 5 <= availableSupplies {
        phase.step = 2
        // phase.purchasedAttempts = numberOfAttempts
        return true
    }
    return false
}

func (phase *ResearchPhase) AttemptResearch(numberOfAttempts int, dice Roller) bool {
    if phase.step != 2 {
        return false
    } else {
        phase.step = 3
        // TODO: Implement this:
        //   Loop for numberOfAttempts
        //   Roll a die, record result
        //   Could then: return results or else award research and return new completed research struct
        return true
    }
    return false
}
