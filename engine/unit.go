package engine

type Unit struct {
    movementRange    int
    attackRating     int
    defenseRating    int
    supplyCost       int
    wounded          bool
    controllingPower string
}

func (u *Unit) Attack(roll int) bool {
    return roll <= u.attackRating
}

func (u *Unit) Defend(roll int) bool {
    return roll <= u.defenseRating
}

func (u *Unit) Wound() bool {
    u.wounded = true
    return true
}

func (u *Unit) Wounded() bool { return u.wounded }

func (u Unit) MovementRange() int { return u.movementRange }

func (u *Unit) ControllingPower() string { return u.controllingPower }

func (u Unit) Side() string {
    return SidesMap[u.controllingPower]
}

type LandUnit struct {
    Unit
}

func (u LandUnit) IsTerrainValid(terrain string) bool {
    if terrain == "land" {
        return true
    }
    return false
}

func (u LandUnit) CanStopInZone(z *Zone) bool {
    return true
}

func (u LandUnit) Category() string { return "land" }

type SeaUnit struct {
    Unit
}

func (u SeaUnit) IsTerrainValid(terrain string) bool {
    if terrain == "sea" {
        return true
    }
    return false
}

func (u SeaUnit) CanStopInZone(z *Zone) bool {
    return true
}

func (u SeaUnit) Category() string { return "sea" }

type AirUnit struct {
    Unit
}

func (u AirUnit) IsTerrainValid(terrain string) bool {
    if terrain == "impassible" {
        return false
    }
    return true
}

func (u AirUnit) CanStopInZone(z *Zone) bool {
    if z.TerrainType() == "sea" {
        return false
    }
    return true
}

func (u AirUnit) Category() string { return "air" }
