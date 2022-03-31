package emath

type Number interface {
	~int | ~int16 | ~int32 | ~uint8 | ~uint16 | ~uint32
}

func Min[N Number](a N, b N) N {
	if a < b {
		return a
	}

	return b
}

func Max[N Number](a N, b N) N {
	if a > b {
		return a
	}

	return b
}

func Clamp[N Number](n N, min N, max N) N {
	return Min(max, Max(min, n))
}
