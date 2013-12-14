package engine

import (
    "time"
    "math/rand"
)

type Dice struct {
    sides int
}

func (d *Dice) Roll() int {
    rand.Seed(time.Now().UnixNano())
    roll := 0
    for roll == 0 {
        roll = rand.Intn(d.sides + 1)
    }
    return roll
}
