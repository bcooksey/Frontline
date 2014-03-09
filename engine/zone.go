package engine

import "container/heap"
import "fmt"

var _ = fmt.Println // TOOD: DELETE

type Zone struct {
    id               int
    occupyingUnits   []Unit
    neighboringZones []*Zone
    supplyValue      int
    terrainType      string
    nativePower      string
    controllingPower string
}

// Getters
func (z *Zone) SupplyValue() int          { return z.supplyValue }
func (z *Zone) OccupyingUnits() []Unit    { return z.occupyingUnits }
func (z *Zone) NeighboringZones() []*Zone { return z.neighboringZones }
func (z *Zone) TerrainType() string       { return z.terrainType }
func (z *Zone) NativePower() string       { return z.nativePower }
func (z *Zone) ControllingPower() string  { return z.controllingPower }

func (z *Zone) AddOccupyingUnit(newUnit Unit) bool {
    z.occupyingUnits = append(z.occupyingUnits, newUnit)
    return true
}

type ZoneQueue []Zone

func (q ZoneQueue) Len() int           { return len(q) }
func (q ZoneQueue) Less(i, j int) bool { panic("Do not run less") }
func (q ZoneQueue) Swap(i, j int)      { q[i], q[j] = q[j], q[i] }

func (q *ZoneQueue) Push(z interface{}) {
    *q = append(*q, z.(Zone))
}

func (q *ZoneQueue) Pop() interface{} {
    old := *q
    n := len(old)
    z := old[n-1]
    *q = old[0 : n-1]
    return z
}

func Move(fromZone Zone, toZone Zone, unit Unit) bool {
    if fromZone.id == toZone.id {
        return true
    }

    queue := &ZoneQueue{}
    visits := map[int]bool{}

    heap.Init(queue)

    visits[fromZone.id] = true
    queue.Push(fromZone)

    movesLeft := unit.MovementRange()

    // Do a breadth-first search to see if the zones connect
    for queue.Len() > 0 {
        if movesLeft <= 0 {
            return false
        }
        currentZone := queue.Pop().(Zone)
        for _, neighbor := range currentZone.NeighboringZones() {
            if !visits[neighbor.id] {
                if !isTerrainValidForUnit(currentZone.terrainType, unit.Category()) {
                    continue
                }

                if neighbor.id == toZone.id {
                    return true
                } else {
                    visits[neighbor.id] = true
                    queue.Push(*neighbor)
                }
            }
        }
        movesLeft--
    }
    return false
}

func isTerrainValidForUnit(terrainType string, unitCategory string) bool {
    if unitCategory == "land" && terrainType == "land" {
        return true
    } else {
        return false
    }
}
