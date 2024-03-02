package emath_test

import (
	"testing"

	. "github.com/epes/emath"
)

func Test_Bytes_Vector2(t *testing.T) {
	tests := []struct {
		v    Vector2
		want []byte
	}{
		{
			v: Vector2{X: 1, Y: 2},
		}, {
			v: Vector2{X: 0, Y: 0},
		}, {
			v: Vector2{X: -1, Y: -2},
		},
	}

	for _, tc := range tests {
		marshaled := tc.v.Bytes()
		unmarshaled := BytesToVector2(marshaled)

		if unmarshaled.Equals(tc.v) == false {
			t.Fatalf("expected: %v, got: %v", tc.v, unmarshaled)
		}
	}
}
