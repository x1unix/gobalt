package mathx

import (
	"math"
)

// Abs returns the absolute value of x.
func Abs[T Number](val T) T {
	return T(math.Abs(float64(val)))
}
