package engine

type Zone struct {
    occupyingUnits   []Unit
    neighboringZones []Zone
    supplyValue      int
    terrainType      string
    nativePower      string
    controllingPower string
}

// Getters
func (z *Zone) SupplyValue() int         { return z.supplyValue }
func (z *Zone) OccupyingUnits() []Unit   { return z.occupyingUnits }
func (z *Zone) NeighboringZones() []Zone { return z.neighboringZones }
func (z *Zone) TerrainType() string      { return z.terrainType }
func (z *Zone) NativePower() string      { return z.nativePower }
func (z *Zone) ControllingPower() string { return z.controllingPower }

func (z *Zone) AddOccupyingUnit(newUnit Unit) bool {
    z.occupyingUnits = append(z.occupyingUnits, newUnit)
    return true
}
