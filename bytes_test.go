package emath_test

import (
	"testing"

	"github.com/epes/emath"
)

func Test_IsNilBytes(t *testing.T) {
	tests := []struct {
		b    []byte
		want bool
	}{
		{
			b:    []byte{0},
			want: true,
		}, {
			b:    []byte{1},
			want: false,
		}, {
			b:    []byte{'0'},
			want: false,
		},
	}

	for _, tc := range tests {
		got := emath.IsNilBytes(tc.b)

		if tc.want != got {
			t.Fatalf("expected: %v, got: %v", tc.want, got)
		}
	}
}

func Fuzz_IsNilBytes(f *testing.F) {
	f.Fuzz(func(t *testing.T, b []byte) {
		emath.IsNilBytes(b)
	})
}
