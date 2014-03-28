package engine

type LandUnit struct {
    Unit
}

func (u LandUnit) IsTerrainValid(terrain string) bool {
    if terrain == "land" {
        return true
    }
    return false
}

func (u LandUnit) CanStopInZone(z Zone) bool {
    return true
}

func (u LandUnit) Category() string { return "land" }
