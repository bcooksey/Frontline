package engine

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
