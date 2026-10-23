package darts

import "math"
func Score(x, y float64) int {
    
    value := (x * x) + (y * y)
    rootValue := math.Sqrt(value)

    if rootValue <= 1 {
        return 10
    } else if rootValue <= 5 {
        return 5
    } else if rootValue <= 10 {
        return 1
    }
    return 0
}