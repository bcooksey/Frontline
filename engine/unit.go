package engine

type Unit struct {
    movementRange    int
    attackRating     int
    defenseRating    int
    supplyCost       int
    wounded          bool
    controllingPower string
}

func (u Unit) Attack(roll int) bool {
    return roll <= u.attackRating
}

func (u Unit) Defend(roll int) bool {
    return roll <= u.defenseRating
}

func (u *Unit) Wound() bool {
    u.wounded = true
    return true
}

func (u Unit) Wounded() bool { return u.wounded }

func (u Unit) MovementRange() int { return u.movementRange }

func (u Unit) ControllingPower() string { return u.controllingPower }

func (u Unit) Side() string {
    return SidesMap[u.controllingPower]
}
