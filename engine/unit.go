package engine

type Unit struct {
    movementRange int
    attackRating  int
    defenseRating int
    supplyCost    int
    category      string
    wounded       bool
}

func (u *Unit) Attack(roll int) bool {
    return roll <= u.attackRating
}

func (u *Unit) Defend(roll int) bool {
    return roll <= u.defenseRating
}

func (u *Unit) Category() string { return u.category }

func (u *Unit) Wound() bool {
    u.wounded = true
    return true
}

func (u *Unit) Wounded() bool { return u.wounded }

func (u *Unit) MovementRange() int { return u.movementRange }
