package engine

/* A Battle involes attacking units and defending uints.
* It consists of rounds, where each round has several phases:
    1. Submarine surprise strike or submerge (sea battles only)
      * Attacking subs roll or submerge
      * Defending subs roll or submerge
      * Defender must select who will die
      * Attacker must select who will die
      * Units die!
    2. Attacking units fire
    3. Defending units fire
    4. Remove defender’s casualties
      * Defender must select who will die
      * Attacker must select who will die
      * Units die!
* After the final phase, one of four things happens:
  1: The attacker decides to attack again - A new round begins
  1: The attacker kills all the defending units - Results in capturing Zone. Battle is over
  2: The attacker withdrawls - Must relocate units back to a friendly Zone. Battle is over
  3: The defender kills all the attackers units - Battle is over
*/

/*
   b = new battle
   b.rollForAttackers()
   b.rollForDefenders()
   b.woundDefenders(units)
   b.woundAttackers(units)
   b.removeCasualties()
*/

type Battle struct {
    attackers []Unit
    defenders []Unit
    phase     string
}

func CreateBattle(attackers []Unit, defenders []Unit) Battle {
    return Battle{attackers: attackers, defenders: defenders, phase: "attack"}
}

// Getters
func (b *Battle) Phase() string { return b.phase }
