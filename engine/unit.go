package engine

type Unit struct {
    movementRange int
    attackRating int
    defenseRating int
    supplyCost int
}

func (u *Unit) Attack(roll int) bool {
    return roll <= u.attackRating
}

func (u *Unit) Defend(roll int) bool {
    return roll <= u.defenseRating
}
