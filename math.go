package emath

type Number interface {
	~int | ~int16 | ~int32 | ~uint8 | ~uint16 | ~uint32 | ~float32 | ~float64
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

func Lerp[N Number](a N, b N, t float64) float64 {
	return float64(a) + float64(b-a)*t
}

func Abs[N Number](n N) N {
	if n < 0 {
		return -n
	}

	return n
}
