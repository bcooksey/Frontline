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

type Roller interface {
    Roll() int
}

type Attacker interface {
    Attack(int) bool
    Category() string
}

type Defender interface {
    Defend(int) bool
    Category() string
}

/*
   b = new battle
   b.rollForAttackers()
   b.rollForDefenders()
   b.woundDefenders(units)
   b.woundAttackers(units)
   b.removeCasualties()
*/

type Battle struct {
    attackers []Attacker
    defenders []Defender
    phase     string
    dice      Roller
}

func CreateBattle(attackers []Attacker, defenders []Defender, dice Roller) Battle {
    if dice == nil {
        dice = &Dice{sides: 6}
    }
    return Battle{
        attackers: attackers,
        defenders: defenders,
        phase:     "attack",
        dice:      dice,
    }
}

// Getters
func (b *Battle) Phase() string { return b.phase }

func (b *Battle) RollForAttackers() map[string]int {
    hits := map[string]int{"land": 0, "sea": 0, "air": 0}
    for _, attacker := range b.attackers {
        if attacker.Attack(b.dice.Roll()) {
            hits[attacker.Category()] += 1
        }
    }
    return hits
}
