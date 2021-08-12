package utils

// MinInt finds min int
//
func MinInt(a int, b ...int) int {
	min := a
	for i := 0; i < len(b); i++ {
		if b[i] < min {
			min = b[i]
		}
	}

	return min
}
