package emath

func Uint8Max(a uint8, b uint8) uint8 {
	if a > b {
		return a
	}

	return b
}

func Uint8Min(a uint8, b uint8) uint8 {
	if a > b {
		return b
	}

	return a
}
