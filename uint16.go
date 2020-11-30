package emath

import (
	"encoding/binary"
)

func Uint16ToBytes(u uint16) []byte {
	b := make([]byte, 2)
	binary.BigEndian.PutUint16(b, u)
	return b
}

func Uint16Max(a uint16, b uint16) uint16 {
	if a > b {
		return a
	}

	return b
}

func Uint16Min(a uint16, b uint16) uint16 {
	if a > b {
		return b
	}

	return a
}
