package mathx

type Reducer[T Number] func(prev, current T) T

// Reduce executes a user-supplied "reducer" callback function on each element of the array,
// in order, passing in the return value from the calculation on the preceding element.
func Reduce[T Number](items []T, reducer Reducer[T]) T {
	if len(items) == 0 {
		return 0
	}

	accumVal := items[0]
	for i := 1; i < len(items); i++ {
		curVal := items[i]
		accumVal = reducer(curVal, accumVal)
	}

	return accumVal
}

// Sum returns summary of all numbers in a slice.
func Sum[T Number](items []T) T {
	return Reduce(items, func(prev, current T) T {
		return prev + current
	})
}

// Avg returns average value of a slice.
func Avg[T Number](items []T) T {
	if len(items) == 0 {
		return 0
	}

	sum := Sum(items)
	return sum / T(len(items))
}

// Convert converts slice of one numeric type into another.
func Convert[B, A Number](src []A) []B {
	result := make([]B, len(src))
	for i, val := range src {
		result[i] = B(val)
	}

	return result
}
