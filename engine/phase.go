package engine

// TODO: Need to decide on interface for a phase. maybe:
//   CurrentStep
//   UpdateTurnFromPhase (for applying new research, purchased units, etc.)
//   NextPhase() - Returns next phase so easy to continue driving through states

import "fmt"
var _ = fmt.Println // TOOD: DELETE

type ResearchPhase struct {
    state ResearchState
    // noAttemptsState ResearchNoAttemptsState
    // attemptsBoughtState ResearchAttemptsBoughtState
    // successState ResearchSuccessState
}

// func (phase *ResearchPhase) Init() {
//    phase.noAttemptsState = &ResearchNoAttemptsState{Phase: *phase}
//    phase.attemptsBoughtState = &ResearchAttemptsBoughtState{Phase: *phase}
//    phase.successState = &ResearchSuccessState{Phase: *phase}
// 
//    phase.state = phase.noAttemptsState
// }

func (phase *ResearchPhase) BuyAttempts(numberOfAttempts int, availableSupplies int) bool {
   return phase.state.BuyAttempts(numberOfAttempts, availableSupplies)
}

func (phase *ResearchPhase) AttemptResearch(numberOfAttempts int, dice Roller) bool {
   return phase.state.AttemptResearch(numberOfAttempts, dice)
}

func (phase *ResearchPhase) AwardSuccesses() bool {
   return phase.state.AwardSuccesses()
}

func (phase ResearchPhase) CurrentState() ResearchState {
   return phase.state
}

func (phase *ResearchPhase) SetState(state ResearchState) {
   phase.state = state
}

type ResearchState interface {
   BuyAttempts(int, int) bool
   AttemptResearch(int, Roller) bool
   AwardSuccesses() bool
}

type ResearchNoAttemptsState struct {
   Phase *ResearchPhase
}

func (state *ResearchNoAttemptsState) BuyAttempts(numberOfAttempts int, availableSupplies int) bool {
   if numberOfAttempts * 5 <= availableSupplies {
      state.Phase.SetState(&ResearchAttemptsBoughtState{Phase: *state.Phase, numberOfAttempts: numberOfAttempts})
      return true
   }
   return false
}

func (state *ResearchNoAttemptsState) AttemptResearch(numberOfAttempts int, dice Roller) bool {
   return false
}

func (state *ResearchNoAttemptsState) AwardSuccesses() bool {
   return false
}

type ResearchAttemptsBoughtState struct {
   Phase ResearchPhase
   numberOfAttempts int
}

func (state ResearchAttemptsBoughtState) GetNumberOfAttempts() int {
   return state.numberOfAttempts
}

func (state ResearchAttemptsBoughtState) BuyAttempts(numberOfAttempts int, availableSupplies int) bool {
   return false
}

func (state ResearchAttemptsBoughtState) AttemptResearch(numberOfAttempts int, dice Roller) bool {
   if numberOfAttempts > 0 {
      // state.Phase.State = ResearchAttemptsBoughtState{Phase: state.Phase}
      // TODO: Implement this:
      //   Loop for numberOfAttempts
      //   Roll a die, record result
      //   Could then: return results or else award research and return new completed research struct
      return true
   }
   return false
}

func (state ResearchAttemptsBoughtState) AwardSuccesses() bool {
   return false
}

type ResearchSuccessState struct {
   Phase ResearchPhase
}

func (state ResearchSuccessState) BuyAttempts(numberOfAttempts int, availableSupplies int) bool {
   return false
}

func (state ResearchSuccessState) AttemptResearch(numberOfAttempts int, dice Roller) bool {
   return false
}

func (state ResearchSuccessState) AwardSuccesses() bool {
   return true
}
