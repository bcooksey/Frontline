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

    unit.movementRange = 2
    if Move(zone1, zone4, unit) != true {
        t.Error("Could not successfully move a unit to a zone in range")
    }
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

    if Move(zone1, zone2, unit) != false {
        t.Error("Land unit stopped in the sea")
    }
}

func TestMoveSeaUnitOnlyGoesOverSea(t *testing.T) {
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

    if Move(zone1, zone2, unit) != false {
        t.Error("Sea unit stopped on land")
    }
}

func TestMoveAirUnitGoesOverAnything(t *testing.T) {
    zone1 := Zone{id: 1, terrainType: "land"}
    zone2 := Zone{id: 2, terrainType: "land"}
    zone3 := Zone{id: 3, terrainType: "land", neighboringZones: []*Zone{&zone2}}
    zone1.neighboringZones = []*Zone{&zone2}
    zone2.neighboringZones = []*Zone{&zone1, &zone3}

    unit := Unit{movementRange: 20, category: "air"}

    if Move(zone1, zone3, unit) != true {
        t.Error("Could not successfully move an air unit from one zone to another")
    }

    zone2.terrainType = "sea"
    if Move(zone1, zone3, unit) != true {
        t.Error("Air unit could not move across sea")
    }

    // Case: Air unit cannot stop in sea
}

func TestMoveImpassibleLandIsImpassible(t *testing.T) {
    zone1 := Zone{id: 1, terrainType: "land"}
    zone2 := Zone{id: 2, terrainType: "impassible"}
    zone3 := Zone{id: 3, terrainType: "land", neighboringZones: []*Zone{&zone2}}
    zone1.neighboringZones = []*Zone{&zone2}
    zone2.neighboringZones = []*Zone{&zone1, &zone3}

    unit := Unit{movementRange: 20, category: "land"}

    if Move(zone1, zone3, unit) != false {
        t.Error("Land unit moved across impassible terrain")
    }

    unit.category = "sea"
    if Move(zone1, zone3, unit) != false {
        t.Error("Sea unit moved across impassible terrain")
    }

    unit.category = "air"
    if Move(zone1, zone3, unit) != false {
        t.Error("Air unit moved across impassible terrain")
    }
}

func TestMoveUnitCanMoveThroughFriendlyZone(t *testing.T) {
    zone1 := Zone{id: 1, terrainType: "land", controllingPower: "us"}
    zone2 := Zone{id: 2, terrainType: "land", controllingPower: "uk"}
    zone3 := Zone{id: 3, terrainType: "land", controllingPower: "us", neighboringZones: []*Zone{&zone2}}
    zone1.neighboringZones = []*Zone{&zone2}
    zone2.neighboringZones = []*Zone{&zone1, &zone3}

    unit := Unit{movementRange: 20, category: "land", controllingPower: "us"}

    if Move(zone1, zone3, unit) != true {
        t.Error("Could not move unit through friendly zone")
    }

    if Move(zone1, zone2, unit) != true {
        t.Error("Could not stop unit in friendly zone")
    }

    zone2.controllingPower = "japan"
    if Move(zone1, zone3, unit) != false {
        t.Error("Unit moved through hostile zone")
    }

    if Move(zone1, zone2, unit) != false {
        t.Error("Unit stopped in hostile zone")
    }

}
