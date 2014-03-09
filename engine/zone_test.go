package engine

import "testing"
import "fmt"

var _ = fmt.Println // TODO: Delete

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

func TestMove(t *testing.T) {
    zone1 := Zone{id: 1, terrainType: "land"}
    zone2 := Zone{id: 2, terrainType: "land", neighboringZones: []*Zone{&zone1}}
    zone3 := Zone{id: 3, terrainType: "land"}
    zone4 := Zone{id: 4, terrainType: "land", neighboringZones: []*Zone{&zone3}}
    zone1.neighboringZones = []*Zone{&zone2, &zone3}
    zone3.neighboringZones = []*Zone{&zone1, &zone4}

    unit := Unit{movementRange: 3, category: "land"}

    if Move(zone1, zone1, unit) != true {
        t.Error("Did not claim keeping a unit in same zone is valid")
    }

    if Move(zone1, zone4, unit) != true {
        t.Error("Could not successfully move a unit from one zone to another")
    }

    /* Cases to test:
     *   Sea units
     *   Air units
     */
}

func TestMoveZoneBeyondUnitMovementRange(t *testing.T) {
    zone1 := Zone{id: 1, terrainType: "land"}
    zone2 := Zone{id: 2, terrainType: "land", neighboringZones: []*Zone{&zone1}}
    zone3 := Zone{id: 3, terrainType: "land"}
    zone4 := Zone{id: 4, terrainType: "land", neighboringZones: []*Zone{&zone3}}
    zone1.neighboringZones = []*Zone{&zone2, &zone3}
    zone3.neighboringZones = []*Zone{&zone1, &zone4}

    unit := Unit{movementRange: 1, category: "land"}

    if Move(zone1, zone4, unit) != false {
        t.Error("Unit moved beyond its movement range")
    }
}

func TestMoveLandUnitOnlyGoesOverLand(t *testing.T) {
    zone1 := Zone{id: 1, terrainType: "land"}
    zone2 := Zone{id: 2, terrainType: "sea"}
    zone3 := Zone{id: 3, terrainType: "land", neighboringZones: []*Zone{&zone2}}
    zone1.neighboringZones = []*Zone{&zone2}
    zone2.neighboringZones = []*Zone{&zone1, &zone3}

    unit := Unit{movementRange: 20, category: "land"}

    if Move(zone1, zone3, unit) != false {
        t.Error("Land unit moved across sea")
    }

    zone2.terrainType = "impassible"
    if Move(zone1, zone3, unit) != false {
        t.Error("Land unit moved across impassible area")
    }
}

func TestMoveSeaUnit(t *testing.T) {
    zone1 := Zone{id: 1, terrainType: "sea"}
    zone2 := Zone{id: 2, terrainType: "sea"}
    zone3 := Zone{id: 3, terrainType: "sea", neighboringZones: []*Zone{&zone2}}
    zone1.neighboringZones = []*Zone{&zone2}
    zone2.neighboringZones = []*Zone{&zone1, &zone3}

    unit := Unit{movementRange: 20, category: "sea"}

    if Move(zone1, zone3, unit) != true {
        t.Error("Could not successfully move a sea unit from one zone to another")
    }

    zone2.terrainType = "land"
    if Move(zone1, zone3, unit) != false {
        t.Error("Sea unit moved across land")
    }
}
