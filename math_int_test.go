package emath_test

import (
	"testing"

	"github.com/epes/emath"
)

func Test_Min_Int(t *testing.T) {
	tests := []struct {
		a    int
		b    int
		want int
	}{
		{
			a:    1,
			b:    2,
			want: 1,
		}, {
			a:    1,
			b:    -1,
			want: -1,
		}, {
			a:    10,
			b:    100,
			want: 10,
		}, {
			a:    0,
			b:    0,
			want: 0,
		},
	}

	for _, tc := range tests {
		got := emath.Min(tc.a, tc.b)

		if tc.want != got {
			t.Fatalf("expected: %v, got: %v", tc.want, got)
		}
	}
}

func Fuzz_Min_Int(f *testing.F) {
	f.Fuzz(func(t *testing.T, a int, b int) {
		emath.Min(a, b)
	})
}

func Test_Max_Int(t *testing.T) {
	tests := []struct {
		a    int
		b    int
		want int
	}{
		{
			a:    1,
			b:    2,
			want: 2,
		}, {
			a:    1,
			b:    -1,
			want: 1,
		}, {
			a:    10,
			b:    100,
			want: 100,
		}, {
			a:    0,
			b:    0,
			want: 0,
		},
	}

	for _, tc := range tests {
		got := emath.Max(tc.a, tc.b)

		if tc.want != got {
			t.Fatalf("expected: %v, got: %v", tc.want, got)
		}
	}
}

func Fuzz_Max_Int(f *testing.F) {
	f.Fuzz(func(t *testing.T, a int, b int) {
		emath.Max(a, b)
	})
}

func Test_Clamp_Int(t *testing.T) {
	tests := []struct {
		n    int
		min  int
		max  int
		want int
	}{
		{
			n:    1,
			min:  0,
			max:  2,
			want: 1,
		}, {
			n:    1,
			min:  0,
			max:  0,
			want: 0,
		}, {
			n:    0,
			min:  0,
			max:  0,
			want: 0,
		}, {
			n:    100,
			min:  0,
			max:  10,
			want: 10,
		}, {
			n:    -100,
			min:  0,
			max:  10,
			want: 0,
		}, {
			n:    int(0),
			min:  int(0),
			max:  int(0),
			want: 0,
		},
	}

	for _, tc := range tests {
		got := emath.Clamp(tc.n, tc.min, tc.max)
		if tc.want != got {
			t.Fatalf("expected: %v, got: %v", tc.want, got)
		}
	}
}

func Fuzz_Clamp_Int(f *testing.F) {
	f.Fuzz(func(t *testing.T, a int, min int, max int) {
		emath.Clamp(a, min, max)
	})
}
