package engine

type SeaUnit struct {
    Unit
}

func (u SeaUnit) IsTerrainValid(terrain string) bool {
    if terrain == "sea" {
        return true
    }
    return false
}

func (u SeaUnit) CanStopInZone(z Zone) bool {
    return true
}

func (u SeaUnit) Category() string { return "sea" }
