package engine

// import "fmt"

/* A Turn consists of several activities (TODO: Verify I didn't miss a step)
   1. Earn supplies
   1. Research
   2. Buy units
   3. Place units at factories that have enough supply production
   4. Choose Battles
   5. Fortify
*/

/* Rules currently being ignored:
 */

type Turn struct {
    power string
    phase string
}
