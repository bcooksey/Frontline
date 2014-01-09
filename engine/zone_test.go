package engine

import "testing"
import "fmt"
var _ = fmt.Println

func TestZoneGetter(t *testing.T) {
    zone := Zone{supplyValue: 2}

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

func TestMoveUnit(t *testing.T) {
    zone1 := Zone{id:1}
    zone2 := Zone{id:2, neighboringZones: []Zone{zone1}}
    zone4 := Zone{id: 4}
    zone3 := Zone{id: 3, neighboringZones: []Zone{zone1, zone4}}
    zone1.neighboringZones = []Zone{zone2, zone3}

    for _, z := range zone1.neighboringZones[1].neighboringZones {
        fmt.Println(z.id)
    }

    unit := Unit{movementRange: 3}

    if Move(zone1, zone1, unit) != true {
        t.Error("Did not claim keeping a unit in same zone is valid")
    }

    if Move(zone1, zone4, unit) != true {
        t.Error("Could not successfully move a unit from one zone to another")
    }

    /* Cases to test:
     *   Zone is beyond the units movement capability
     *   One path is beyond units movement, another is ok
     *   Land unit trying to go through sea
     *   Land unit having to go through impassible area
     *   Sea units
     *   Air units
     */
}
