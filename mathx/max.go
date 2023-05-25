package mathx

// Max returns the larger of x or y.
func Max[T Number](x, y T) T {
	if x > y {
		return x
	}

	return y
}

// Min returns the smaller of x or y.
func Min[T Number](x, y T) T {
	if x < y {
		return x
	}

	return y
}

// MaxElement returns the largest item in slice.
func MaxElement[T Number](items []T) T {
	return Reduce(items, Max[T])
}

// MinElement returns the smallest item in slice.
func MinElement[T Number](items []T) T {
	return Reduce(items, Min[T])
}
