package math

// MinInt returns minimum of two numbers
func MinInt(a, b int) int {
	if a < b {
		return a
	}
	return b
}

// MinUint returns minimum of two numbers
func MinUint(a, b uint) uint {
	if a < b {
		return a
	}
	return b
}

// MaxInt returns minimum of two numbers
func MaxInt(a, b int) int {
	if a > b {
		return a
	}
	return b
}

// MaxUint returns minimum of two numbers
func MaxUint(a, b uint) uint {
	if a > b {
		return a
	}
	return b
}

// BoundInt checks the int against boundary
func BoundInt(val, min, max int) int {
	switch {
	case val > max:
		return max
	case val < min:
		return min
	default:
		return val
	}
}

// BoundUnt checks the int against boundary
func BoundUnt(val, min, max uint) uint {
	switch {
	case val > max:
		return max
	case val < min:
		return min
	default:
		return val
	}
}

// CeilAfterDiv Returns ceiling value after dividing
func CeilAfterDiv(num, denom int) int {
	if num%denom == 0 {
		return num / denom
	}
	return (num / denom) + 1
}
