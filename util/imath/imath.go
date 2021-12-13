package imath

// Abs calculates the absolute number of an int
func Abs(x int) int {
	if x < 0 {
		return -x
	}
	return x
}

// Min computes the minimum value
func Min(head int, xs ...int) int {
	min := head
	for _, x := range xs {
		if x < min {
			min = x
		}
	}
	return min
}

// Max computes the maximum value
func Max(head int, xs ...int) int {
	max := head
	for _, x := range xs {
		if x > max {
			max = x
		}
	}
	return max
}

// Clamp clamps the integer x to [-1, 1], the result is 0 if x is zero
func Clamp(x int) int {
	switch {
	case x > 0:
		return 1
	case x < 0:
		return -1
	default:
		return 0
	}
}
