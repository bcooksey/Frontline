package main

type Unit struct {
    movementRange int
    attackRating int
    defenseRating int 
    supplyCost int
}

func (u *Unit) attack(roll int) bool {
    return roll <= u.attackRating
}

func (u *Unit) defend(roll int) bool {
    return roll <= u.defenseRating
}
