package engine

import "fmt"

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

/* Rules currently being ignored:
   Special combat phases (bombing, sea support)
   Units being able to take multiple hits
*/

type Roller interface {
    Roll() int
}

type Soldier interface {
    Wound() bool
    Wounded() bool
    Category() string
}

type Attacker interface {
    Attack(int) bool
    Soldier
}

type Defender interface {
    Defend(int) bool
    Soldier
}

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
func (b *Battle) Phase() string         { return b.phase }
func (b *Battle) Attackers() []Attacker { return b.attackers }
func (b *Battle) Defenders() []Defender { return b.defenders }

func (b *Battle) RollForAttackers() map[string]int {
    hits := map[string]int{"land": 0, "sea": 0, "air": 0}
    for _, attacker := range b.attackers {
        if attacker.Attack(b.dice.Roll()) {
            hits[attacker.Category()] += 1
        }
    }
    return hits
}

func (b *Battle) RollForDefenders() map[string]int {
    hits := map[string]int{"land": 0, "sea": 0, "air": 0}
    for _, defender := range b.defenders {
        if defender.Defend(b.dice.Roll()) {
            hits[defender.Category()] += 1
        }
    }
    return hits
}

func (b *Battle) woundSoldiers(units []Soldier, casualties map[string]int) (bool, error) {
    // After the sort, map contains pointers to battle's units
    sorted_units := map[string][]Soldier{}
    for _, unit := range units {
        sorted_units[unit.Category()] = append(sorted_units[unit.Category()], unit)
    }

    for category, count := range casualties {
        if count > len(sorted_units[category]) {
            return false, IllegalOperationError{
                message: fmt.Sprintf("Removing %d %s units failed. Only %d available.", count, category, len(sorted_units[category])),
            }
        }
        for i := 0; i < count; i++ {
            sorted_units[category][i].Wound()
        }
    }
    return true, nil
}

func (b *Battle) WoundDefenders(casualties map[string]int) (bool, error) {
    var units []Soldier
    for _, defender := range b.defenders {
        units = append(units, defender)
    }
    return b.woundSoldiers(units, casualties)
}

func (b *Battle) WoundAttackers(casualties map[string]int) (bool, error) {
    var units []Soldier
    for _, attacker := range b.attackers {
        units = append(units, attacker)
    }
    return b.woundSoldiers(units, casualties)
}

func (b *Battle) RemoveCasualties() bool {
    return true
}

type IllegalOperationError struct {
    message string
}

func (e IllegalOperationError) Error() string {
    return fmt.Sprintf("%s", e.message)
}
