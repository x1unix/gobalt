package mathx

import "golang.org/x/exp/constraints"

// Number is a constraint that represents any numeric value, except complex numbers.
type Number interface {
	constraints.Float | constraints.Integer
}
