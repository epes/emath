package emath_test

import (
	"testing"

	"github.com/epes/emath"
)

func Test_Lerp_Float64(t *testing.T) {
	tests := []struct {
		a    float64
		b    float64
		t    float64
		want float64
	}{
		{0, 1, 0.5, 0.5},
		{1, 2, 0.5, 1.5},
		{0, 3, 1.0 / 3.0, 1},
		{1, 2, 0, 1},
		{1, 2, 1, 2},
		{-1, 1, 0.5, 0},
		{0, 0, 0.5, 0},
		{0, 0, 0, 0},
		{0, 0, 1, 0},
	}

	for _, tc := range tests {
		got := emath.Lerp(tc.a, tc.b, tc.t)
		if tc.want != got {
			t.Fatalf("expected: %v, got: %v", tc.want, got)
		}
	}
}

func Fuzz_Lerp_Float64(f *testing.F) {
	f.Fuzz(func(tt *testing.T, a float64, b float64, t float64) {
		emath.Lerp(a, b, t)
	})
}

func Test_Abs_Float64(t *testing.T) {
	tests := []struct {
		n    float64
		want float64
	}{
		{1, 1},
		{-1, 1},
		{0, 0},
		{1.5, 1.5},
		{-1.5, 1.5},
		{0.5, 0.5},
		{-0.5, 0.5},
		{0.0, 0.0},
	}

	for _, tc := range tests {
		got := emath.Abs(tc.n)
		if tc.want != got {
			t.Fatalf("expected: %v, got: %v", tc.want, got)
		}
	}
}

func Fuzz_Abs_Float64(f *testing.F) {
	f.Fuzz(func(tt *testing.T, n float64) {
		emath.Abs(n)
	})
}
