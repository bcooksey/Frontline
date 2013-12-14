package engine

import "testing"

func TestZoneGetter(t *testing.T) {
    zone := Zone{supplyValue:  2}

    if zone.SupplyValue() != 2 {
        t.Error("Zone getter broken")
    }
}

func TestAddUnitToZone(t *testing.T) {
    zone := Zone{}
    unit := Unit{}

    zone.AddOccupyingUnit(unit)

    if len(zone.OccupyingUnits()) != 1 {
        t.Error("AddOccupyingUnit did not add unit to zone")
    }
}
